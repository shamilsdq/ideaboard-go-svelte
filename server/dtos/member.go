package dtos

type MemberJoinBroadcastDto struct {
	MemberCount int `json:"memberCount"`
}

type MemberExitBroadcastDto = MemberJoinBroadcastDto
