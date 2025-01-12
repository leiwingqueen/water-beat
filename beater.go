package main

import "sort"

type Beater struct {
	matchList []match
}

// 队伍
type team struct {
	member1 string
	member2 string
}

// 比赛记录
type match struct {
	team1  team
	team2  team
	score1 int
	score2 int
	water  int
}

// 打水结果
type betResult struct {
	from  string
	to    string
	water int
}

// 计算打水结果
func (b *Beater) calculate() []betResult {
	// 统计每个人的赢和输的水
	mp := make(map[string]int)
	matchList := b.matchList
	for _, item := range matchList {
		if item.score1 > item.score2 {
			mp[item.team1.member1] += item.water
			mp[item.team1.member2] += item.water
			mp[item.team2.member1] -= item.water
			mp[item.team2.member2] -= item.water
		} else {
			mp[item.team1.member1] -= item.water
			mp[item.team1.member2] -= item.water
			mp[item.team2.member1] += item.water
			mp[item.team2.member2] += item.water
		}
	}
	var userList []string
	for user := range mp {
		userList = append(userList, user)
	}
	sort.Slice(userList, func(i, j int) bool {
		return mp[userList[i]] < mp[userList[j]]
	})
	// 双指针计算每个人要付的水
	l, r := 0, len(userList)-1
	var res []betResult
	for l < r {
		user1 := userList[l]
		user2 := userList[r]
		wl := mp[user1]
		wr := mp[user2]
		if wl == 0 {
			l++
			continue
		}
		if wr == 0 {
			r--
			continue
		}
		w := min(abs(wl), abs(wr))
		res = append(res, betResult{user1, user2, w})
		wl += w
		wr -= w
		mp[user1] = wl
		mp[user2] = wr
		if wl == 0 {
			l++
		}
		if wr == 0 {
			r--
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}
}
