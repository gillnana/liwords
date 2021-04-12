import React from 'react';

import { useCallback, useEffect, useMemo } from 'react';

import { useParams } from 'react-router-dom';
import axios from 'axios';
import { message } from 'antd';
import { clubRedirects } from '../lobby/fixed_seek_controls';
import {
  useLoginStateStoreContext,
  useTournamentStoreContext,
} from '../store/store';
import { useMountedState } from '../utils/mounted';
import { postBinary, toAPIUrl } from '../api/api';
import { TopBar } from '../topbar/topbar';
import { singularCount } from '../utils/plural';
import { Chat } from '../chat/chat';
import { TournamentInfo } from './tournament_info';
import { sendAccept, sendSeek } from '../lobby/sought_game_interactions';
import { SoughtGame } from '../store/reducers/lobby_reducer';
import { ActionsPanel } from './actions_panel';
import { CompetitorStatus } from './competitor_status';
import { ActionType } from '../actions/actions';
import { readyForTournamentGame } from '../store/reducers/tournament_reducer';
import './room.scss';
import {
  GetTournamentMetadataRequest,
  TournamentMetadataResponse,
} from '../gen/api/proto/tournament_service/tournament_service_pb';

type Props = {
  sendSocketMsg: (msg: Uint8Array) => void;
  sendChat: (msg: string, chan: string) => void;
};

export const TournamentRoom = (props: Props) => {
  const { useState } = useMountedState();

  const { partialSlug } = useParams();
  const { loginState } = useLoginStateStoreContext();
  const {
    tournamentContext,
    dispatchTournamentContext,
  } = useTournamentStoreContext();
  const { loggedIn, username, userID } = loginState;
  const { competitorState: competitorContext } = tournamentContext;
  const { isRegistered } = competitorContext;
  const { sendSocketMsg } = props;
  const { path } = loginState;
  const [badTournament, setBadTournament] = useState(false);
  const [selectedGameTab, setSelectedGameTab] = useState('GAMES');

  useEffect(() => {
    if (!partialSlug || !path) {
      return;
    }
    // Temporary redirect code:
    if (path.startsWith('/tournament/')) {
      const oldslug = path.substr('/tournament/'.length);
      if (oldslug in clubRedirects) {
        const slug = clubRedirects[oldslug];
        window.location.replace(
          `${window.location.protocol}//${window.location.hostname}${slug}`
        );
      }
    }
    const req = new GetTournamentMetadataRequest();
    req.setSlug(path);

    postBinary(
      'tournament_service.TournamentService',
      'GetTournamentMetadata',
      req
    )
      .then((rbin) => {
        const resp = TournamentMetadataResponse.deserializeBinary(
          rbin.data
        ).toObject();
        console.log('resp', resp);
        dispatchTournamentContext({
          actionType: ActionType.SetTourneyMetadata,
          payload: resp,
        });
      })
      .catch((err) => {
        message.error({
          content: 'Error fetching tournament data',
          duration: 5,
        });
        setBadTournament(true);
      });
  }, [path, partialSlug, dispatchTournamentContext]);

  const tournamentID = useMemo(() => {
    return tournamentContext.metadata.id;
  }, [tournamentContext.metadata]);

  // Should be more like "amdirector"
  const isDirector = useMemo(() => {
    console.log('tournament metadata', tournamentContext.metadata);
    return tournamentContext.metadata.directorsList.includes(username);
  }, [tournamentContext.metadata, username]);

  const handleNewGame = useCallback(
    (seekID: string) => {
      sendAccept(seekID, sendSocketMsg);
    },
    [sendSocketMsg]
  );
  const onSeekSubmit = useCallback(
    (g: SoughtGame) => {
      sendSeek(g, sendSocketMsg);
    },
    [sendSocketMsg]
  );

  const peopleOnlineContext = useCallback(
    (n: number) => singularCount(n, 'Player', 'Players'),
    []
  );

  if (badTournament) {
    return (
      <>
        <TopBar />
        <div className="lobby">
          <h3>You tried to access a non-existing page.</h3>
        </div>
      </>
    );
  }

  if (!tournamentID) {
    return (
      <>
        <TopBar />
      </>
    );
  }

  return (
    <>
      <TopBar />
      <div className={`lobby room ${isRegistered ? ' competitor' : ''}`}>
        <div className="chat-area">
          <Chat
            sendChat={props.sendChat}
            defaultChannel={`chat.tournament.${tournamentID}`}
            defaultDescription={tournamentContext.metadata.name}
            peopleOnlineContext={peopleOnlineContext}
            highlight={tournamentContext.metadata.directorsList}
            highlightText="Director"
            tournamentID={tournamentID}
          />
          {isRegistered && (
            <CompetitorStatus
              sendReady={() =>
                readyForTournamentGame(
                  sendSocketMsg,
                  tournamentContext.metadata.id,
                  competitorContext
                )
              }
            />
          )}
        </div>
        <ActionsPanel
          selectedGameTab={selectedGameTab}
          setSelectedGameTab={setSelectedGameTab}
          isDirector={isDirector}
          tournamentID={tournamentID}
          onSeekSubmit={onSeekSubmit}
          loggedIn={loggedIn}
          newGame={handleNewGame}
          username={username}
          userID={userID}
          sendReady={() =>
            readyForTournamentGame(
              sendSocketMsg,
              tournamentContext.metadata.id,
              competitorContext
            )
          }
        />
        <TournamentInfo
          setSelectedGameTab={setSelectedGameTab}
          sendSocketMsg={sendSocketMsg}
        />
      </div>
    </>
  );
};
