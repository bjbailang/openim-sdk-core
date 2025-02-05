// Copyright © 2023 OpenIM SDK. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sdk_params_callback

import (
	"github.com/openimsdk/openim-sdk-core/v3/sdk_struct"
)

type ConversationArgs struct {
	ConversationID  string   `json:"conversationID"`
	ClientMsgIDList []string `json:"clientMsgIDList"`
}

type FindMessageListCallback struct {
	TotalCount      int                           `json:"totalCount"`
	FindResultItems []*SearchByConversationResult `json:"findResultItems"`
}

type GetAdvancedHistoryMessageListParams struct {
	ConversationID   string `json:"conversationID"`
	StartClientMsgID string `json:"startClientMsgID"`
	Count            int    `json:"count"`
	ViewType         int    `json:"viewType"`
}

type GetAdvancedHistoryMessageListCallback struct {
	MessageList []*sdk_struct.MsgStruct `json:"messageList"`
	IsEnd       bool                    `json:"isEnd"`
	ErrCode     int32                   `json:"errCode"`
	ErrMsg      string                  `json:"errMsg"`
}

type FetchSurroundingMessagesReq struct {
	StartMessage *sdk_struct.MsgStruct `json:"startMessage"`
	ViewType     int                   `json:"viewType"`
	Before       int                   `json:"before"`
	After        int                   `json:"after"`
}
type FetchSurroundingMessagesResp struct {
	MessageList []*sdk_struct.MsgStruct `json:"messageList"`
}

type SearchLocalMessagesParams struct {
	ConversationID       string   `json:"conversationID"`
	KeywordList          []string `json:"keywordList"`
	KeywordListMatchType int      `json:"keywordListMatchType"`
	SenderUserIDList     []string `json:"senderUserIDList"`
	MessageTypeList      []int    `json:"messageTypeList"`
	SearchTimePosition   int64    `json:"searchTimePosition"`
	SearchTimePeriod     int64    `json:"searchTimePeriod"`
	PageIndex            int      `json:"pageIndex"`
	Count                int      `json:"count"`
}

type SearchLocalMessagesCallback struct {
	TotalCount        int                           `json:"totalCount"`
	SearchResultItems []*SearchByConversationResult `json:"searchResultItems"`
}

type SearchByConversationResult struct {
	ConversationID    string                  `json:"conversationID"`
	ConversationType  int32                   `json:"conversationType"`
	ShowName          string                  `json:"showName"`
	FaceURL           string                  `json:"faceURL"`
	LatestMsgSendTime int64                   `json:"latestMsgSendTime,omitempty"`
	MessageCount      int                     `json:"messageCount"`
	MessageList       []*sdk_struct.MsgStruct `json:"messageList"`
}
