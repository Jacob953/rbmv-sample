package rbmv

import (
	"math"
)

func NewReBalanceRes(shard *Shard, from, to string) *ReBalanceRes {
	return &ReBalanceRes{
		shard: shard,
		from:  from,
		to:    to,
	}
}

type ReBalanceRes struct {
	shard *Shard
	from  string
	to    string
}

func ReBalance(saveNodes, delNodes []*Node) (res *ReBalanceRes, exist bool) {
	currV := math.MaxFloat64
	saveMemList, memAvg := initMem(saveNodes, delNodes)
	for i := range delNodes {
		if !delNodes[i].HasShards() {
			continue
		}
		shard := delNodes[i].Pop()
		for j := range saveNodes {
			if saveNodes[j].ExistShard(shard.name) {
				continue
			}
			saveMemList[j] += shard.mem
			nextV := Variance(saveMemList, memAvg)
			if hasBetterMove(nextV, currV) {
				currV = nextV
				res = NewReBalanceRes(shard, delNodes[i].peer, saveNodes[j].peer)
				exist = true
			}
			saveMemList[j] -= shard.mem
		}
		if exist {
			return res, true
		}
	}
	return nil, false
}

func initMem(saveNodes, delNodes []*Node) ([]float64, float64) {
	memSum := 0.0
	saveMemList := make([]float64, len(saveNodes))
	for i := range saveNodes {
		memSum += saveNodes[i].mem
		saveMemList = append(saveMemList, saveNodes[i].mem)
	}
	for i := range delNodes {
		memSum += delNodes[i].mem
	}
	memAvg := memSum / float64(len(saveNodes)+len(delNodes))
	return saveMemList, memAvg
}

func hasBetterMove(nextV, currV float64) bool {
	return nextV < currV
}
