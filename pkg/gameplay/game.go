// Package game should know nothing about protocols or databases.
// It is mostly a pass-through interface to a Macondo game,
// but also implements a timer and other related logic.
// This is a use-case in the clean architecture hierarchy.
package gameplay

import (
	"context"
	"errors"

	"github.com/domino14/macondo/alphabet"
	"github.com/rs/zerolog/log"

	"github.com/domino14/macondo/move"

	macondopb "github.com/domino14/macondo/gen/api/proto/macondo"

	"github.com/domino14/liwords/pkg/config"
	"github.com/domino14/liwords/pkg/entity"
	pb "github.com/domino14/liwords/rpc/api/proto"
	"github.com/domino14/macondo/board"
	"github.com/domino14/macondo/game"
)

const (
	CrosswordGame string = "CrosswordGame"
)

var (
	errGameNotActive   = errors.New("game is not currently active")
	errNotOnTurn       = errors.New("player not on turn")
	errTimeDidntRunOut = errors.New("got time ran out, but it did not actually")
)

// GameStore is an interface for getting a full game.
type GameStore interface {
	Get(ctx context.Context, id string) (*entity.Game, error)
	Set(context.Context, *entity.Game) error
}

// type Game interface {
// }

// InstantiateNewGame instantiates a game and returns it.
func InstantiateNewGame(ctx context.Context, gameStore GameStore, cfg *config.Config,
	players []*macondopb.PlayerInfo, req *pb.GameRequest) (*entity.Game, error) {

	var bd []string
	if req.Rules == nil {
		return nil, errors.New("no rules")
	}
	switch req.Rules.BoardLayoutName {
	case CrosswordGame:
		bd = board.CrosswordGameBoard
	default:
		return nil, errors.New("unsupported board layout")
	}

	rules, err := game.NewGameRules(&cfg.MacondoConfig, bd,
		req.Lexicon, req.Rules.LetterDistributionName)

	if err != nil {
		return nil, err
	}
	g, err := game.NewGame(rules, players)
	if err != nil {
		return nil, err
	}
	// StartGame creates a new history Uid and deals tiles, etc.
	g.StartGame()
	g.SetBackupMode(game.InteractiveGameplayMode)
	g.SetChallengeRule(req.ChallengeRule)

	entGame := entity.NewGame(g, req)
	// Save the game to the store.
	if err = gameStore.Set(ctx, entGame); err != nil {
		return nil, err
	}
	return entGame, nil
	// We return the instantiated game. Although the tiles have technically been
	// dealt out, we need to call StartGame to actually start the timer
	// and forward game events to the right channels.
}

// func StartGameInstance(entGame *entity.Game, eventChan chan<- *entity.EventWrapper) error {
// 	if err := entGame.RegisterChangeHook(eventChan); err != nil {
// 		return err
// 	}
// 	entGame.
// 	entGame.SendChange(entity.WrapEvent(entGame.HistoryRefresherEvent(), pb.MessageType_GAME_HISTORY_REFRESHER,
// 		entGame.GameID()))

// 	return nil
// }

func clientEventToMove(cge *pb.ClientGameplayEvent, g *game.Game) (*move.Move, error) {
	playerid := g.PlayerOnTurn()
	rack := g.RackFor(playerid)

	switch cge.Type {
	case pb.ClientGameplayEvent_TILE_PLACEMENT:
		m, err := g.CreateAndScorePlacementMove(cge.PositionCoords, cge.Tiles, rack.String())
		if err != nil {
			return nil, err
		}
		log.Debug().Msg("got a client gameplay event tile placement")
		// Note that we don't validate the move here, but we do so later.
		return m, nil

	case pb.ClientGameplayEvent_PASS:
		m := move.NewPassMove(rack.TilesOn(), g.Alphabet())
		return m, nil
	case pb.ClientGameplayEvent_EXCHANGE:
		tiles, err := alphabet.ToMachineWord(cge.Tiles, g.Alphabet())
		if err != nil {
			return nil, err
		}
		leaveMW, err := game.Leave(rack.TilesOn(), tiles)
		if err != nil {
			return nil, err
		}
		m := move.NewExchangeMove(tiles, leaveMW, g.Alphabet())
		return m, nil
	}
	return nil, errors.New("client gameplay event not handled")
}

func StartGame(ctx context.Context, gameStore GameStore, eventChan chan<- *entity.EventWrapper, id string) error {
	// Note that StartGame does _not_ start the Macondo game, which
	// has already started, but we don't "know" that. It is _this_
	// function that will actually start the game in the user's eyes.
	// It needs to reset the timer to now.
	entGame, err := gameStore.Get(ctx, id)
	if err != nil {
		return err
	}
	// This should be True, see comment above.
	if entGame.Game.Playing() != macondopb.PlayState_PLAYING {
		return errGameNotActive
	}
	log.Debug().Str("gameid", id).Msg("reset timers (and start)")
	entGame.ResetTimers()

	// Save the game back to the store always.
	if err := gameStore.Set(ctx, entGame); err != nil {
		return err
	}
	if err := entGame.RegisterChangeHook(eventChan); err != nil {
		return err
	}
	log.Debug().Interface("history", entGame.Game.History()).Msg("game history")

	// We do not send a history refresher event here. Instead, we will send one
	// when the user joins the game realm. See the `sendRealmData` function.

	return nil
}

func handleChallenge(ctx context.Context, entGame *entity.Game, gameStore GameStore) error {
	if entGame.ChallengeRule() == macondopb.ChallengeRule_VOID {
		// The front-end shouldn't even show the button.
		return errors.New("challenges not acceptable in void")
	}
	challenger := entGame.Game.NickOnTurn()

	valid, err := entGame.Game.ChallengeEvent(0)
	if err != nil {
		return err
	}
	resultEvent := &pb.ServerChallengeResultEvent{
		Valid:         valid,
		ChallengeRule: entGame.ChallengeRule(),
		Challenger:    challenger,
	}
	entGame.SendChange(entity.WrapEvent(resultEvent, pb.MessageType_SERVER_CHALLENGE_RESULT_EVENT,
		entGame.GameID()))

	// Send a refresher event to get the game state up-to-date on the client.
	entGame.SendChange(entity.WrapEvent(entGame.HistoryRefresherEvent(), pb.MessageType_GAME_HISTORY_REFRESHER,
		entGame.GameID()))

	if entGame.Game.Playing() == macondopb.PlayState_GAME_OVER {
		performEndgameDuties(entGame)
	}

	err = gameStore.Set(ctx, entGame)
	if err != nil {
		return err
	}

	return nil
}

func PlayMove(ctx context.Context, gameStore GameStore, player string,
	cge *pb.ClientGameplayEvent) error {

	// XXX: VERIFY THAT THE CLIENT GAME ID CORRESPONDS TO THE GAME
	// THE PLAYER IS PLAYING!
	entGame, err := gameStore.Get(ctx, cge.GameId)
	if err != nil {
		return err
	}
	if entGame.Game.Playing() == macondopb.PlayState_GAME_OVER {
		return errGameNotActive
	}
	onTurn := entGame.Game.PlayerOnTurn()

	// Ensure that it is actually the correct player's turn
	if entGame.Game.NickOnTurn() != player {
		return errNotOnTurn
	}

	// Check that we didn't run out of time.
	if entGame.TimeRemaining(onTurn) < 0 {
		// Game is over!
		entGame.Game.SetPlaying(macondopb.PlayState_GAME_OVER)
		// Basically skip to the bottom and exit.
		return setTimedOut(ctx, entGame, gameStore)
	}

	log.Debug().Msg("going to turn into a macondo gameevent")

	if cge.Type == pb.ClientGameplayEvent_CHALLENGE_PLAY {
		// Handle in another way
		return handleChallenge(ctx, entGame, gameStore)
	}

	// Turn the event into a macondo GameEvent.
	m, err := clientEventToMove(cge, &entGame.Game)
	if err != nil {
		return err
	}
	// m := game.MoveFromEvent(evt, entGame.Game.Alphabet(), entGame.Game.Board())
	log.Debug().Msg("validating")

	_, err = entGame.Game.ValidateMove(m)
	if err != nil {
		return err
	}

	// Don't back up the move, but add to history
	log.Debug().Msg("playing the move")
	err = entGame.Game.PlayMove(m, true)
	if err != nil {
		return err
	}
	// Get the turn that we _just_ appended to the history
	turnLength := len(entGame.Game.History().Turns)
	turn := entGame.Game.History().Turns[turnLength-1]
	// Register time.
	entGame.RecordTimeOfMove(onTurn)

	// Create a set of ServerGameplayEvents to send back.

	evts := make([]*pb.ServerGameplayEvent, len(turn.Events))

	for idx, evt := range turn.Events {

		sge := &pb.ServerGameplayEvent{}
		sge.Event = evt
		sge.GameId = cge.GameId
		// note that `onTurn` is correct as it was saved up there before
		// we played the turn.
		sge.TimeRemaining = int32(entGame.TimeRemaining(onTurn))
		sge.NewRack = entGame.Game.RackLettersFor(onTurn)
		sge.Playing = entGame.Game.Playing()
		evts[idx] = sge
	}
	// Since the move was successful, we assume the user gameplay event is valid.
	// Re-send it, but overwrite the time remaining and new rack properly.
	playing := entGame.Game.Playing()

	err = gameStore.Set(ctx, entGame)
	if err != nil {
		return err
	}
	for _, sge := range evts {
		entGame.SendChange(entity.WrapEvent(sge, pb.MessageType_SERVER_GAMEPLAY_EVENT,
			entGame.GameID()))
	}
	if playing == macondopb.PlayState_GAME_OVER {
		performEndgameDuties(entGame)
	}
	return nil
}

func TimedOut(ctx context.Context, gameStore GameStore, player string, gameID string) error {
	// XXX: VERIFY THAT THE GAME ID is the client's current game!!
	entGame, err := gameStore.Get(ctx, gameID)
	if err != nil {
		return err
	}
	if entGame.Game.Playing() == macondopb.PlayState_GAME_OVER {
		return errGameNotActive
	}
	onTurn := entGame.Game.PlayerOnTurn()

	// Ensure that it is actually the correct player's turn
	if entGame.Game.NickOnTurn() != player {
		return errNotOnTurn
	}
	if entGame.TimeRemaining(onTurn) > 0 {
		return errTimeDidntRunOut
	}
	// Ok, the time did run out after all.

	return setTimedOut(ctx, entGame, gameStore)
}

func setTimedOut(ctx context.Context, entGame *entity.Game, gameStore GameStore) error {
	log.Debug().Msg("timed out!")
	entGame.Game.SetPlaying(macondopb.PlayState_GAME_OVER)
	// Store the game back into the store
	err := gameStore.Set(ctx, entGame)
	if err != nil {
		return err
	}
	// And send a game end event.
	performEndgameDuties(entGame)
	return nil
}

func performEndgameDuties(g *entity.Game) {
	// XXX: need to log who lost or won, the timeout state, etc.
	g.SendChange(
		entity.WrapEvent(g.GameEndedEvent(),
			pb.MessageType_GAME_ENDED_EVENT, g.GameID()))

}
