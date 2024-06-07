package foodgroup

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/mk6i/retro-aim-server/config"

	"github.com/stretchr/testify/assert"

	"github.com/mk6i/retro-aim-server/state"
	"github.com/mk6i/retro-aim-server/wire"
)

func TestOServiceServiceForBOS_ServiceRequest(t *testing.T) {
	cases := []struct {
		// name is the unit test name
		name string
		// config is the application config
		cfg config.Config
		// chatRoom is the chat room the user connects to
		chatRoom *state.ChatRoom
		// userSession is the session of the user requesting the chat service
		// info
		userSession *state.Session
		// inputSNAC is the SNAC sent by the sender client
		inputSNAC wire.SNACMessage
		// expectSNACFrame is the SNAC frame sent from the server to the recipient
		// client
		expectOutput wire.SNACMessage
		// mockParams is the list of params sent to mocks that satisfy this
		// method's dependencies
		mockParams mockParams
		// expectErr is the expected error returned by the router
		expectErr error
	}{
		{
			name:        "request info for ICBM service, return invalid SNAC err",
			userSession: newTestSession("user_screen_name"),
			inputSNAC: wire.SNACMessage{
				Frame: wire.SNACFrame{
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x04_OServiceServiceRequest{
					FoodGroup: wire.ICBM,
				},
			},
			expectOutput: wire.SNACMessage{
				Frame: wire.SNACFrame{
					FoodGroup: wire.OService,
					SubGroup:  wire.OServiceErr,
					RequestID: 1234,
				},
				Body: wire.SNACError{
					Code: wire.ErrorCodeServiceUnavailable,
				},
			},
		},
		{
			name: "request info for connecting to chat nav, return chat nav connection metadata",
			cfg: config.Config{
				OSCARHost:   "127.0.0.1",
				ChatNavPort: "1234",
			},
			userSession: newTestSession("user_screen_name"),
			inputSNAC: wire.SNACMessage{
				Frame: wire.SNACFrame{
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x04_OServiceServiceRequest{
					FoodGroup: wire.ChatNav,
				},
			},
			expectOutput: wire.SNACMessage{
				Frame: wire.SNACFrame{
					FoodGroup: wire.OService,
					SubGroup:  wire.OServiceServiceResponse,
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x05_OServiceServiceResponse{
					TLVRestBlock: wire.TLVRestBlock{
						TLVList: wire.TLVList{
							wire.NewTLV(wire.OServiceTLVTagsReconnectHere, "127.0.0.1:1234"),
							wire.NewTLV(wire.OServiceTLVTagsLoginCookie, []byte("the-cookie")),
							wire.NewTLV(wire.OServiceTLVTagsGroupID, wire.ChatNav),
							wire.NewTLV(wire.OServiceTLVTagsSSLCertName, ""),
							wire.NewTLV(wire.OServiceTLVTagsSSLState, uint8(0x00)),
						},
					},
				},
			},
			mockParams: mockParams{
				cookieIssuerParams: cookieIssuerParams{
					{
						data:   []byte("user_screen_name"),
						cookie: []byte("the-cookie"),
					},
				},
			},
		},
		{
			name: "request info for connecting to alert svc, return alert svc connection metadata",
			cfg: config.Config{
				OSCARHost: "127.0.0.1",
				AlertPort: "1234",
			},
			userSession: newTestSession("user_screen_name"),
			inputSNAC: wire.SNACMessage{
				Frame: wire.SNACFrame{
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x04_OServiceServiceRequest{
					FoodGroup: wire.Alert,
				},
			},
			expectOutput: wire.SNACMessage{
				Frame: wire.SNACFrame{
					FoodGroup: wire.OService,
					SubGroup:  wire.OServiceServiceResponse,
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x05_OServiceServiceResponse{
					TLVRestBlock: wire.TLVRestBlock{
						TLVList: wire.TLVList{
							wire.NewTLV(wire.OServiceTLVTagsReconnectHere, "127.0.0.1:1234"),
							wire.NewTLV(wire.OServiceTLVTagsLoginCookie, []byte("the-cookie")),
							wire.NewTLV(wire.OServiceTLVTagsGroupID, wire.Alert),
							wire.NewTLV(wire.OServiceTLVTagsSSLCertName, ""),
							wire.NewTLV(wire.OServiceTLVTagsSSLState, uint8(0x00)),
						},
					},
				},
			},
			mockParams: mockParams{
				cookieIssuerParams: cookieIssuerParams{
					{
						data:   []byte("user_screen_name"),
						cookie: []byte("the-cookie"),
					},
				},
			},
		},
		{
			name: "request info for connecting to chat room, return chat service and chat room metadata",
			cfg: config.Config{
				OSCARHost: "127.0.0.1",
				ChatPort:  "1234",
			},
			chatRoom: &state.ChatRoom{
				CreateTime:     time.UnixMilli(0),
				DetailLevel:    4,
				Exchange:       8,
				Cookie:         "the-chat-cookie",
				InstanceNumber: 16,
				Name:           "my new chat",
			},
			userSession: newTestSession("user_screen_name"),
			inputSNAC: wire.SNACMessage{
				Frame: wire.SNACFrame{
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x04_OServiceServiceRequest{
					FoodGroup: wire.Chat,
					TLVRestBlock: wire.TLVRestBlock{
						TLVList: wire.TLVList{
							wire.NewTLV(0x01, wire.SNAC_0x01_0x04_TLVRoomInfo{
								Exchange:       8,
								Cookie:         "the-chat-cookie",
								InstanceNumber: 16,
							}),
						},
					},
				},
			},
			expectOutput: wire.SNACMessage{
				Frame: wire.SNACFrame{
					FoodGroup: wire.OService,
					SubGroup:  wire.OServiceServiceResponse,
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x05_OServiceServiceResponse{
					TLVRestBlock: wire.TLVRestBlock{
						TLVList: wire.TLVList{
							wire.NewTLV(wire.OServiceTLVTagsReconnectHere, "127.0.0.1:1234"),
							wire.NewTLV(wire.OServiceTLVTagsLoginCookie, []byte("the-cookie")),
							wire.NewTLV(wire.OServiceTLVTagsGroupID, wire.Chat),
							wire.NewTLV(wire.OServiceTLVTagsSSLCertName, ""),
							wire.NewTLV(wire.OServiceTLVTagsSSLState, uint8(0x00)),
						},
					},
				},
			},
			mockParams: mockParams{
				cookieIssuerParams: cookieIssuerParams{
					{
						data: []byte{
							0x0F, 't', 'h', 'e', '-', 'c', 'h', 'a', 't', '-', 'c', 'o', 'o', 'k', 'i', 'e',
							0x10, 'u', 's', 'e', 'r', '_', 's', 'c', 'r', 'e', 'e', 'n', '_', 'n', 'a', 'm', 'e',
						},
						cookie: []byte("the-cookie"),
					},
				},
			},
		},
		{
			name: "request info for connecting to non-existent chat room, return ErrChatRoomNotFound",
			cfg: config.Config{
				OSCARHost: "127.0.0.1",
				ChatPort:  "1234",
			},
			chatRoom:    nil,
			userSession: newTestSession("user_screen_name"),
			inputSNAC: wire.SNACMessage{
				Frame: wire.SNACFrame{
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x04_OServiceServiceRequest{
					FoodGroup: wire.Chat,
					TLVRestBlock: wire.TLVRestBlock{
						TLVList: wire.TLVList{
							wire.NewTLV(0x01, wire.SNAC_0x01_0x04_TLVRoomInfo{
								Exchange:       8,
								Cookie:         "the-chat-cookie",
								InstanceNumber: 16,
							}),
						},
					},
				},
			},
			expectErr: state.ErrChatRoomNotFound,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			//
			// initialize dependencies
			//
			sessionManager := newMockSessionManager(t)
			chatRegistry := state.NewChatRegistry()
			chatSess := &state.Session{}
			if tc.chatRoom != nil {
				sessionManager.EXPECT().
					AddSession(tc.userSession.ScreenName()).
					Return(chatSess).
					Maybe()
				chatRegistry.Register(*tc.chatRoom, sessionManager)
			}
			cookieIssuer := newMockCookieIssuer(t)
			for _, params := range tc.mockParams.cookieIssuerParams {
				cookieIssuer.EXPECT().
					Issue(params.data).
					Return(params.cookie, params.err)
			}
			//
			// send input SNAC
			//
			svc := NewOServiceServiceForBOS(tc.cfg, nil, nil, slog.Default(), cookieIssuer, nil, chatRegistry)

			outputSNAC, err := svc.ServiceRequest(nil, tc.userSession, tc.inputSNAC.Frame,
				tc.inputSNAC.Body.(wire.SNAC_0x01_0x04_OServiceServiceRequest))
			assert.ErrorIs(t, err, tc.expectErr)
			if tc.expectErr != nil {
				return
			}
			//
			// verify output
			//
			assert.Equal(t, tc.expectOutput, outputSNAC)
		})
	}
}

func TestSetUserInfoFields(t *testing.T) {
	cases := []struct {
		// name is the unit test name
		name string
		// userSession is the session of the user whose info is being set
		userSession *state.Session
		// inputSNAC is the SNAC sent from the client to the server
		inputSNAC wire.SNACMessage
		// expectOutput is the SNAC reply sent from the server back to the
		// client
		expectOutput wire.SNACMessage
		// broadcastMessage is the arrival/departure message sent to buddies
		broadcastMessage []struct {
			recipients []string
			msg        wire.SNACMessage
		}
		// interestedUserLookups contains all the users who have this user on
		// their buddy list
		interestedUserLookups map[string][]string
		// expectErr is the expected error returned
		expectErr error
		// mockParams is the list of params sent to mocks that satisfy this
		// method's dependencies
		mockParams mockParams
	}{
		{
			name:        "set user status to visible",
			userSession: newTestSession("user_screen_name"),
			inputSNAC: wire.SNACMessage{
				Frame: wire.SNACFrame{
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields{
					TLVRestBlock: wire.TLVRestBlock{
						TLVList: wire.TLVList{
							wire.NewTLV(wire.OServiceUserInfoStatus, uint32(0x0000)),
						},
					},
				},
			},
			expectOutput: wire.SNACMessage{
				Frame: wire.SNACFrame{
					FoodGroup: wire.OService,
					SubGroup:  wire.OServiceUserInfoUpdate,
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x0F_OServiceUserInfoUpdate{
					TLVUserInfo: newTestSession("user_screen_name").TLVUserInfo(),
				},
			},
			mockParams: mockParams{
				buddyBroadcasterParams: buddyBroadcasterParams{
					broadcastBuddyArrivedParams: broadcastBuddyArrivedParams{
						{
							screenName: "user_screen_name",
						},
					},
				},
			},
		},
		{
			name:        "set user status to invisible",
			userSession: newTestSession("user_screen_name"),
			inputSNAC: wire.SNACMessage{
				Frame: wire.SNACFrame{
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields{
					TLVRestBlock: wire.TLVRestBlock{
						TLVList: wire.TLVList{
							wire.NewTLV(wire.OServiceUserInfoStatus, uint32(0x0100)),
						},
					},
				},
			},
			expectOutput: wire.SNACMessage{
				Frame: wire.SNACFrame{
					FoodGroup: wire.OService,
					SubGroup:  wire.OServiceUserInfoUpdate,
					RequestID: 1234,
				},
				Body: wire.SNAC_0x01_0x0F_OServiceUserInfoUpdate{
					TLVUserInfo: newTestSession("user_screen_name", sessOptInvisible).TLVUserInfo(),
				},
			},
			mockParams: mockParams{
				buddyBroadcasterParams: buddyBroadcasterParams{
					broadcastBuddyDepartedParams: broadcastBuddyDepartedParams{
						{
							screenName: "user_screen_name",
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			buddyUpdateBroadcaster := newMockBuddyBroadcaster(t)
			for _, params := range tc.mockParams.broadcastBuddyArrivedParams {
				p := params
				buddyUpdateBroadcaster.EXPECT().
					BroadcastBuddyArrived(mock.Anything, mock.MatchedBy(func(s *state.Session) bool {
						return s.ScreenName() == p.screenName
					})).
					Return(nil)
			}
			for _, params := range tc.mockParams.broadcastBuddyDepartedParams {
				p := params
				buddyUpdateBroadcaster.EXPECT().
					BroadcastBuddyDeparted(mock.Anything, mock.MatchedBy(func(s *state.Session) bool {
						return s.ScreenName() == p.screenName
					})).
					Return(nil)
			}
			svc := OServiceService{
				cfg:                    config.Config{},
				logger:                 slog.Default(),
				buddyUpdateBroadcaster: buddyUpdateBroadcaster,
			}
			outputSNAC, err := svc.SetUserInfoFields(nil, tc.userSession, tc.inputSNAC.Frame,
				tc.inputSNAC.Body.(wire.SNAC_0x01_0x1E_OServiceSetUserInfoFields))
			assert.ErrorIs(t, err, tc.expectErr)
			if tc.expectErr != nil {
				return
			}
			assert.Equal(t, tc.expectOutput, outputSNAC)
		})
	}
}

func TestOServiceService_RateParamsQuery(t *testing.T) {
	svc := OServiceService{
		cfg:    config.Config{},
		logger: slog.Default(),
	}

	have := svc.RateParamsQuery(nil, wire.SNACFrame{RequestID: 1234})
	want := wire.SNACMessage{
		Frame: wire.SNACFrame{
			FoodGroup: wire.OService,
			SubGroup:  wire.OServiceRateParamsReply,
			RequestID: 1234,
		},
		Body: wire.SNAC_0x01_0x07_OServiceRateParamsReply{
			RateClasses: []struct {
				ID              uint16
				WindowSize      uint32
				ClearLevel      uint32
				AlertLevel      uint32
				LimitLevel      uint32
				DisconnectLevel uint32
				CurrentLevel    uint32
				MaxLevel        uint32
				LastTime        uint32
				CurrentState    uint8
			}{
				{
					ID:              0x0001,
					WindowSize:      0x00000050,
					ClearLevel:      0x000009C4,
					AlertLevel:      0x000007D0,
					LimitLevel:      0x000005DC,
					DisconnectLevel: 0x00000320,
					CurrentLevel:    0x00000D69,
					MaxLevel:        0x00001770,
					LastTime:        0x00000000,
					CurrentState:    0x00,
				},
			},
			RateGroups: []struct {
				ID    uint16
				Pairs []struct {
					FoodGroup uint16
					SubGroup  uint16
				} `count_prefix:"uint16"`
			}{
				{
					ID: 1,
					Pairs: []struct {
						FoodGroup uint16
						SubGroup  uint16
					}{
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceErr,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceClientOnline,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceHostOnline,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceServiceRequest,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceServiceResponse,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceRateParamsQuery,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceRateParamsReply,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceRateParamsSubAdd,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceRateDelParamSub,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceRateParamChange,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServicePauseReq,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServicePauseAck,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceResume,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceUserInfoQuery,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceUserInfoUpdate,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceEvilNotification,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceIdleNotification,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceMigrateGroups,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceMotd,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceSetPrivacyFlags,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceWellKnownUrls,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceNoop,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceClientVersions,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceHostVersions,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceMaxConfigQuery,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceMaxConfigReply,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceStoreConfig,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceConfigQuery,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceConfigReply,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceSetUserInfoFields,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceProbeReq,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceProbeAck,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceBartReply,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceBartQuery2,
						},
						{
							FoodGroup: wire.OService,
							SubGroup:  wire.OServiceBartReply2,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateErr,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateRightsQuery,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateRightsReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateSetInfo,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateUserInfoQuery,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateUserInfoReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateWatcherSubRequest,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateWatcherNotification,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateSetDirInfo,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateSetDirReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateGetDirInfo,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateGetDirReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateGroupCapabilityQuery,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateGroupCapabilityReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateSetKeywordInfo,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateSetKeywordReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateGetKeywordInfo,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateGetKeywordReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateFindListByEmail,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateFindListReply,
						},
						{
							FoodGroup: wire.Locate,
							SubGroup:  wire.LocateUserInfoQuery2,
						},

						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyErr,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyRightsQuery,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyRightsReply,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyAddBuddies,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyDelBuddies,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyWatcherListQuery,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyWatcherListResponse,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyWatcherSubRequest,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyWatcherNotification,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyRejectNotification,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyArrived,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyDeparted,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyAddTempBuddies,
						},
						{
							FoodGroup: wire.Buddy,
							SubGroup:  wire.BuddyDelTempBuddies,
						},

						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMErr,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMAddParameters,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMDelParameters,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMParameterQuery,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMParameterReply,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMChannelMsgToHost,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMChannelMsgToClient,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMEvilRequest,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMEvilReply,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMMissedCalls,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMClientErr,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMHostAck,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMSinStored,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMSinListQuery,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMSinListReply,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMSinRetrieve,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMSinDelete,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMNotifyRequest,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMNotifyReply,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMClientEvent,
						},
						{
							FoodGroup: wire.ICBM,
							SubGroup:  wire.ICBMSinReply,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavErr,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavRequestChatRights,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavRequestExchangeInfo,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavRequestRoomInfo,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavRequestMoreRoomInfo,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavRequestOccupantList,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavSearchForRoom,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavCreateRoom,
						},
						{
							FoodGroup: wire.ChatNav,
							SubGroup:  wire.ChatNavNavInfo,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatErr,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatRoomInfoUpdate,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUsersJoined,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUsersLeft,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatChannelMsgToHost,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatChannelMsgToClient,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatEvilRequest,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatEvilReply,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatClientErr,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatPauseRoomReq,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatPauseRoomAck,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatResumeRoom,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatShowMyRow,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatShowRowByUsername,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatShowRowByNumber,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatShowRowByName,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatRowInfo,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatListRows,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatRowListInfo,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatMoreRows,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatMoveToRow,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatToggleChat,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatSendQuestion,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatSendComment,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatTallyVote,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatAcceptBid,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatSendInvite,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatDeclineInvite,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatAcceptInvite,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatNotifyMessage,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatGotoRow,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatStageUserJoin,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatStageUserLeft,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUnnamedSnac22,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatClose,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUserBan,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUserUnban,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatJoined,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUnnamedSnac27,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUnnamedSnac28,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUnnamedSnac29,
						},
						{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatRoomInfoOwner,
						},

						{
							FoodGroup: wire.BART,
							SubGroup:  wire.BARTErr,
						},
						{
							FoodGroup: wire.BART,
							SubGroup:  wire.BARTUploadQuery,
						},
						{
							FoodGroup: wire.BART,
							SubGroup:  wire.BARTUploadReply,
						},
						{
							FoodGroup: wire.BART,
							SubGroup:  wire.BARTDownloadQuery,
						},
						{
							FoodGroup: wire.BART,
							SubGroup:  wire.BARTDownloadReply,
						},
						{
							FoodGroup: wire.BART,
							SubGroup:  wire.BARTDownload2Query,
						},
						{
							FoodGroup: wire.BART,
							SubGroup:  wire.BARTDownload2Reply,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagErr,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRightsQuery,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRightsReply,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagQuery,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagQueryIfModified,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagReply,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagUse,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagInsertItem,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagUpdateItem,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagDeleteItem,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagInsertClass,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagUpdateClass,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagDeleteClass,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagStatus,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagReplyNotModified,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagDeleteUser,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagStartCluster,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagEndCluster,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagAuthorizeBuddy,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagPreAuthorizeBuddy,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagPreAuthorizedBuddy,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRemoveMe,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRemoveMe2,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRequestAuthorizeToHost,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRequestAuthorizeToClient,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRespondAuthorizeToHost,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRespondAuthorizeToClient,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagBuddyAdded,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRequestAuthorizeToBadog,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRespondAuthorizeToBadog,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagBuddyAddedToBadog,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagTestSnac,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagForwardMsg,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagIsAuthRequiredQuery,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagIsAuthRequiredReply,
						},
						{
							FoodGroup: wire.Feedbag,
							SubGroup:  wire.FeedbagRecentBuddyUpdate,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPErr,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPLoginRequest,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPLoginResponse,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPRegisterRequest,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPChallengeRequest,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPChallengeResponse,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPAsasnRequest,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPSecuridRequest,
						},
						{
							FoodGroup: wire.BUCP,
							SubGroup:  wire.BUCPRegistrationImageRequest,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertErr,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertSetAlertRequest,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertSetAlertReply,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertGetSubsRequest,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertGetSubsResponse,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertNotifyCapabilities,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertNotify,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertGetRuleRequest,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertGetRuleReply,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertGetFeedRequest,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertGetFeedReply,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertRefreshFeed,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertEvent,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertQogSnac,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertRefreshFeedStock,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertNotifyTransport,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertSetAlertRequestV2,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertSetAlertReplyV2,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertTransitReply,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertNotifyAck,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertNotifyDisplayCapabilities,
						},
						{
							FoodGroup: wire.Alert,
							SubGroup:  wire.AlertUserOnline,
						},
					},
				},
			},
		},
	}

	assert.Equal(t, want, have)
}

func TestOServiceServiceForBOS_OServiceHostOnline(t *testing.T) {
	cookieIssuer := newMockCookieIssuer(t)
	svc := NewOServiceServiceForBOS(config.Config{}, nil, nil, slog.Default(), cookieIssuer, nil, nil)

	want := wire.SNACMessage{
		Frame: wire.SNACFrame{
			FoodGroup: wire.OService,
			SubGroup:  wire.OServiceHostOnline,
		},
		Body: wire.SNAC_0x01_0x03_OServiceHostOnline{
			FoodGroups: []uint16{
				wire.Alert,
				wire.Buddy,
				wire.ChatNav,
				wire.Feedbag,
				wire.ICBM,
				wire.Locate,
				wire.OService,
				wire.BART,
				wire.PermitDeny,
			},
		},
	}

	have := svc.HostOnline()
	assert.Equal(t, want, have)
}

func TestOServiceServiceForChat_OServiceHostOnline(t *testing.T) {
	svc := NewOServiceServiceForChat(config.Config{}, slog.Default(), nil, nil)

	want := wire.SNACMessage{
		Frame: wire.SNACFrame{
			FoodGroup: wire.OService,
			SubGroup:  wire.OServiceHostOnline,
		},
		Body: wire.SNAC_0x01_0x03_OServiceHostOnline{
			FoodGroups: []uint16{
				wire.OService,
				wire.Chat,
			},
		},
	}

	have := svc.HostOnline()
	assert.Equal(t, want, have)
}

func TestOServiceService_ClientVersions(t *testing.T) {
	svc := OServiceService{
		cfg:    config.Config{},
		logger: slog.Default(),
	}

	want := wire.SNACMessage{
		Frame: wire.SNACFrame{
			FoodGroup: wire.OService,
			SubGroup:  wire.OServiceHostVersions,
			RequestID: 1234,
		},
		Body: wire.SNAC_0x01_0x18_OServiceHostVersions{
			Versions: []uint16{5, 6, 7, 8},
		},
	}

	have := svc.ClientVersions(nil, wire.SNACFrame{
		RequestID: 1234,
	}, wire.SNAC_0x01_0x17_OServiceClientVersions{
		Versions: []uint16{5, 6, 7, 8},
	})

	assert.Equal(t, want, have)
}

func TestOServiceService_UserInfoQuery(t *testing.T) {
	svc := OServiceService{
		cfg:    config.Config{},
		logger: slog.Default(),
	}
	sess := newTestSession("test-user")

	want := wire.SNACMessage{
		Frame: wire.SNACFrame{
			FoodGroup: wire.OService,
			SubGroup:  wire.OServiceUserInfoUpdate,
			RequestID: 1234,
		},
		Body: wire.SNAC_0x01_0x0F_OServiceUserInfoUpdate{
			TLVUserInfo: sess.TLVUserInfo(),
		},
	}

	have := svc.UserInfoQuery(nil, sess, wire.SNACFrame{RequestID: 1234})

	assert.Equal(t, want, have)
}

func TestOServiceService_IdleNotification(t *testing.T) {
	tests := []struct {
		name   string
		sess   *state.Session
		bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification
		// mockParams is the list of params sent to mocks that satisfy this
		// method's dependencies
		mockParams mockParams
		wantErr    error
	}{
		{
			name: "set idle from active",
			sess: newTestSession("test-user"),
			bodyIn: wire.SNAC_0x01_0x11_OServiceIdleNotification{
				IdleTime: 90,
			},
			mockParams: mockParams{
				buddyBroadcasterParams: buddyBroadcasterParams{
					broadcastBuddyArrivedParams: broadcastBuddyArrivedParams{
						{
							screenName: "test-user",
						},
					},
				},
			},
		},
		{
			name: "set active from idle",
			sess: newTestSession("test-user", sessOptIdle(90*time.Second)),
			bodyIn: wire.SNAC_0x01_0x11_OServiceIdleNotification{
				IdleTime: 0,
			},
			mockParams: mockParams{
				buddyBroadcasterParams: buddyBroadcasterParams{
					broadcastBuddyArrivedParams: broadcastBuddyArrivedParams{
						{
							screenName: "test-user",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buddyUpdateBroadcaster := newMockBuddyBroadcaster(t)
			for _, params := range tt.mockParams.broadcastBuddyArrivedParams {
				p := params
				buddyUpdateBroadcaster.EXPECT().
					BroadcastBuddyArrived(mock.Anything, mock.MatchedBy(func(s *state.Session) bool {
						return s.ScreenName() == p.screenName
					})).
					Return(nil)
			}
			svc := OServiceService{
				cfg:                    config.Config{},
				logger:                 slog.Default(),
				buddyUpdateBroadcaster: buddyUpdateBroadcaster,
			}
			haveErr := svc.IdleNotification(nil, tt.sess, tt.bodyIn)
			assert.ErrorIs(t, tt.wantErr, haveErr)
		})
	}
}

func TestOServiceServiceForBOS_ClientOnline(t *testing.T) {
	tests := []struct {
		// name is the name of the test
		name string
		// joiningChatter is the session of the arriving user
		sess *state.Session
		// bodyIn is the SNAC body sent from the arriving user's client to the
		// server
		bodyIn wire.SNAC_0x01_0x02_OServiceClientOnline
		// retrieveByScreenNameParams contains params for looking up the
		// session for each of the arriving user's buddies
		retrieveByScreenNameParams retrieveByScreenNameParams
		wantErr                    error
		// mockParams is the list of params sent to mocks that satisfy this
		// method's dependencies
		mockParams mockParams
	}{
		{
			name:   "notify feedbag + client-side buddies that user is online, populate client-side buddy list",
			sess:   newTestSession("test-user"),
			bodyIn: wire.SNAC_0x01_0x02_OServiceClientOnline{},
			mockParams: mockParams{
				legacyBuddyListManagerParams: legacyBuddyListManagerParams{
					legacyBuddiesParams: legacyBuddiesParams{
						{
							userScreenName: "test-user",
							result:         []string{"buddy1", "buddy2"},
						},
					},
				},
				messageRelayerParams: messageRelayerParams{
					retrieveByScreenNameParams: retrieveByScreenNameParams{
						{
							screenName: "buddy1",
							sess:       newTestSession("buddy1", sessOptCannedSignonTime),
						},
						{
							screenName: "buddy2",
							sess:       newTestSession("buddy2", sessOptCannedSignonTime),
						},
					},
				},
				buddyBroadcasterParams: buddyBroadcasterParams{
					broadcastBuddyArrivedParams: broadcastBuddyArrivedParams{
						{
							screenName: "test-user",
						},
					},
					unicastBuddyArrivedParams: unicastBuddyArrivedParams{
						{
							from: "buddy1",
							to:   "test-user",
						},
						{
							from: "buddy2",
							to:   "test-user",
						},
					},
				},
			},
		},
		{
			name:   "notify feedbag + client-side buddies that user is online, don't send invisible buddy arrival",
			sess:   newTestSession("test-user"),
			bodyIn: wire.SNAC_0x01_0x02_OServiceClientOnline{},
			mockParams: mockParams{
				legacyBuddyListManagerParams: legacyBuddyListManagerParams{
					legacyBuddiesParams: legacyBuddiesParams{
						{
							userScreenName: "test-user",
							result:         []string{"buddy1", "buddy2"},
						},
					},
				},
				messageRelayerParams: messageRelayerParams{
					retrieveByScreenNameParams: retrieveByScreenNameParams{
						{
							screenName: "buddy1",
							sess:       newTestSession("buddy1", sessOptCannedSignonTime, sessOptInvisible),
						},
						{
							screenName: "buddy2",
							sess:       newTestSession("buddy2", sessOptCannedSignonTime),
						},
					},
				},
				buddyBroadcasterParams: buddyBroadcasterParams{
					broadcastBuddyArrivedParams: broadcastBuddyArrivedParams{
						{
							screenName: "test-user",
						},
					},
					unicastBuddyArrivedParams: unicastBuddyArrivedParams{
						{
							from: "buddy2",
							to:   "test-user",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			messageRelayer := newMockMessageRelayer(t)
			for _, params := range tt.mockParams.retrieveByScreenNameParams {
				messageRelayer.EXPECT().
					RetrieveByScreenName(params.screenName).
					Return(params.sess)
			}
			legacyBuddyListManager := newMockLegacyBuddyListManager(t)
			for _, params := range tt.mockParams.legacyBuddiesParams {
				legacyBuddyListManager.EXPECT().
					Buddies(params.userScreenName).
					Return(params.result)
			}
			buddyUpdateBroadcaster := newMockBuddyBroadcaster(t)
			for _, params := range tt.mockParams.broadcastBuddyArrivedParams {
				p := params
				buddyUpdateBroadcaster.EXPECT().
					BroadcastBuddyArrived(mock.Anything, mock.MatchedBy(func(s *state.Session) bool {
						return s.ScreenName() == p.screenName
					})).
					Return(nil)
			}
			for _, params := range tt.mockParams.unicastBuddyArrivedParams {
				p := params
				buddyUpdateBroadcaster.EXPECT().
					UnicastBuddyArrived(mock.Anything,
						mock.MatchedBy(func(s *state.Session) bool {
							return s.ScreenName() == p.from
						}),
						mock.MatchedBy(func(s *state.Session) bool {
							return s.ScreenName() == p.to
						})).
					Return(p.err)
			}

			svc := NewOServiceServiceForBOS(config.Config{}, messageRelayer, legacyBuddyListManager, slog.Default(), nil, buddyUpdateBroadcaster, nil)
			haveErr := svc.ClientOnline(nil, tt.bodyIn, tt.sess)
			assert.ErrorIs(t, tt.wantErr, haveErr)
		})
	}
}

func TestOServiceServiceForChat_ClientOnline(t *testing.T) {
	chatter1 := newTestSession("chatter-1", sessOptChatRoomCookie("the-cookie"))
	chatter2 := newTestSession("chatter-2", sessOptChatRoomCookie("the-cookie"))
	chatRoom := state.ChatRoom{
		Cookie:         "the-cookie",
		DetailLevel:    1,
		Exchange:       2,
		InstanceNumber: 3,
		Name:           "the-chat-room",
	}

	type participantsParams []*state.Session
	type broadcastExcept []struct {
		sess    *state.Session
		message wire.SNACMessage
	}
	type sendToScreenNameParams []struct {
		screenName string
		message    wire.SNACMessage
	}

	tests := []struct {
		// name is the name of the test
		name string
		// joiningChatter is the user joining the chat room
		joiningChatter *state.Session
		// bodyIn is the SNAC body sent from the arriving user's client to the
		// server
		bodyIn wire.SNAC_0x01_0x02_OServiceClientOnline
		// participantsParams contains all the chat room participants
		participantsParams participantsParams
		// broadcastExcept contains params for broadcasting chat arrival to all
		// chat participants except the user joining
		broadcastExcept broadcastExcept
		// relayToScreenNameParams contains params for sending chat room
		// metadata and chat participant list to joining user
		sendToScreenNameParams sendToScreenNameParams
		wantErr                error
		// mockParams is the list of params sent to mocks that satisfy this
		// method's dependencies
		mockParams mockParams
	}{
		{
			name:           "upon joining, send chat room metadata and participant list to joining user; alert arrival to existing participants",
			joiningChatter: chatter1,
			bodyIn:         wire.SNAC_0x01_0x02_OServiceClientOnline{},
			broadcastExcept: broadcastExcept{
				{
					sess: chatter1,
					message: wire.SNACMessage{
						Frame: wire.SNACFrame{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUsersJoined,
						},
						Body: wire.SNAC_0x0E_0x03_ChatUsersJoined{
							Users: []wire.TLVUserInfo{
								chatter1.TLVUserInfo(),
							},
						},
					},
				},
			},
			participantsParams: participantsParams{
				chatter1,
				chatter2,
			},
			sendToScreenNameParams: sendToScreenNameParams{
				{
					screenName: chatter1.ScreenName(),
					message: wire.SNACMessage{
						Frame: wire.SNACFrame{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatRoomInfoUpdate,
						},
						Body: wire.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
							Exchange:       chatRoom.Exchange,
							Cookie:         chatRoom.Cookie,
							InstanceNumber: chatRoom.InstanceNumber,
							DetailLevel:    chatRoom.DetailLevel,
							TLVBlock: wire.TLVBlock{
								TLVList: chatRoom.TLVList(),
							},
						},
					},
				},
				{
					screenName: chatter1.ScreenName(),
					message: wire.SNACMessage{
						Frame: wire.SNACFrame{
							FoodGroup: wire.Chat,
							SubGroup:  wire.ChatUsersJoined,
						},
						Body: wire.SNAC_0x0E_0x03_ChatUsersJoined{
							Users: []wire.TLVUserInfo{
								chatter1.TLVUserInfo(),
								chatter2.TLVUserInfo(),
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chatMessageRelayer := newMockChatMessageRelayer(t)
			for _, params := range tt.broadcastExcept {
				chatMessageRelayer.EXPECT().
					RelayToAllExcept(mock.Anything, params.sess, params.message).
					Maybe()
			}
			chatMessageRelayer.EXPECT().
				AllSessions().
				Return(tt.participantsParams).
				Maybe()
			for _, params := range tt.sendToScreenNameParams {
				chatMessageRelayer.EXPECT().
					RelayToScreenName(mock.Anything, params.screenName, params.message).
					Maybe()
			}

			chatRegistry := state.NewChatRegistry()
			chatRegistry.Register(chatRoom, chatMessageRelayer)

			svc := NewOServiceServiceForChat(config.Config{}, slog.Default(), nil, chatRegistry)

			haveErr := svc.ClientOnline(nil, wire.SNAC_0x01_0x02_OServiceClientOnline{}, tt.joiningChatter)
			assert.ErrorIs(t, tt.wantErr, haveErr)
		})
	}
}

func TestOServiceServiceForChatNav_HostOnline(t *testing.T) {
	svc := NewOServiceServiceForChatNav(config.Config{}, slog.Default(), nil)

	want := wire.SNACMessage{
		Frame: wire.SNACFrame{
			FoodGroup: wire.OService,
			SubGroup:  wire.OServiceHostOnline,
		},
		Body: wire.SNAC_0x01_0x03_OServiceHostOnline{
			FoodGroups: []uint16{
				wire.ChatNav,
				wire.OService,
			},
		},
	}

	have := svc.HostOnline()
	assert.Equal(t, want, have)
}

func TestOServiceServiceForAlert_HostOnline(t *testing.T) {
	svc := NewOServiceServiceForAlert(config.Config{}, slog.Default(), nil)

	want := wire.SNACMessage{
		Frame: wire.SNACFrame{
			FoodGroup: wire.OService,
			SubGroup:  wire.OServiceHostOnline,
		},
		Body: wire.SNAC_0x01_0x03_OServiceHostOnline{
			FoodGroups: []uint16{
				wire.Alert,
				wire.OService,
			},
		},
	}

	have := svc.HostOnline()
	assert.Equal(t, want, have)
}

func TestOServiceService_SetPrivacyFlags(t *testing.T) {
	svc := OServiceService{
		cfg:    config.Config{},
		logger: slog.Default(),
	}
	body := wire.SNAC_0x01_0x14_OServiceSetPrivacyFlags{
		PrivacyFlags: wire.OServicePrivacyFlagMember | wire.OServicePrivacyFlagIdle,
	}
	svc.SetPrivacyFlags(context.Background(), body)
}
