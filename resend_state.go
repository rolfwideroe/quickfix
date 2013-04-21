package quickfixgo

type resendState struct {
	inSession
}

func (state resendState) FixMsgIn(session *session, msg Message) (nextState sessionState) {
	for ok := true; ok; {
		nextState = state.inSession.FixMsgIn(session, msg)

		switch nextState.(type) {
		case logoutState, latentState, resendState:
			return
		}

		msg, ok = session.messageStash[session.expectedSeqNum]
	}

	if len(session.messageStash) != 0 {
		nextState = resendState{}
	}

	return
}
