package proposal_manager

import (
	"github.com/devfans/zion-sdk/contracts/native/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
	"main/base"
	"math/big"
)

type ProposalType uint8
type Status uint8

const (
	Normal              ProposalType = 0
	UpdateGlobalConfig  ProposalType = 1
	UpdateCommunityInfo ProposalType = 2

	NOTPASS Status = 0
	PASS    Status = 1
	FAIL    Status = 2

	ProposalListLen int = 20
)

type ProposalList struct {
	ProposalList []*big.Int
}

func (m *ProposalList) Decode(payload []byte) error {
	var data struct {
		ProposalList []byte
	}
	if err := utils.UnpackOutputs(ABI, base.MethodGetProposalList, &data, payload); err != nil {
		return err
	}
	return rlp.DecodeBytes(data.ProposalList, m)
}

type ConfigProposalList struct {
	ConfigProposalList []*big.Int
}

func (m *ConfigProposalList) Decode(payload []byte) error {
	var data struct {
		ProposalList []byte
	}
	if err := utils.UnpackOutputs(ABI, base.MethodGetConfigProposalList, &data, payload); err != nil {
		return err
	}
	return rlp.DecodeBytes(data.ProposalList, m)
}

type CommunityProposalList struct {
	CommunityProposalList []*big.Int
}

func (m *CommunityProposalList) Decode(payload []byte) error {
	var data struct {
		ProposalList []byte
	}
	if err := utils.UnpackOutputs(ABI, base.MethodGetCommunityProposalList, &data, payload); err != nil {
		return err
	}
	return rlp.DecodeBytes(data.ProposalList, m)
}

type Proposal struct {
	ID        *big.Int
	Address   common.Address
	Type      ProposalType
	Content   []byte
	EndHeight *big.Int
	Stake     *big.Int
	Status    Status
}

func (m *Proposal) Decode(payload []byte) error {
	var data struct {
		Proposal []byte
	}
	if err := utils.UnpackOutputs(ABI, base.MethodGetProposal, &data, payload); err != nil {
		return err
	}
	return rlp.DecodeBytes(data.Proposal, m)
}
