package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type Role int

const (
	Default Role = 0
	Trading Role = 1
	CoLead  Role = 2
	Lead    Role = 3
)

func (role *Role) String() string {

	switch *role {
	case 0:
		return "Member"
	case 1:
		return "Trading"
	case 2:
		return "CoLead"
	case 3:
		return "Lead"
	default:
		return "UnCategorized"
	}

}

type Member struct {
	id   int      `json:"Id"`
	role Role     `json:"Role"`
	lan  Language `json:"Language"`
}

func (mem *Member) String() string {
	return fmt.Sprintf("Id: %d\nRole: %s\nLan: %s", mem.id, mem.role.String(), mem.lan.String())
}

// func (mem *Member) promote(role Role) {

// }

type ListMem []*Member
type MembersData struct {
	Members []Member `json:"members"`
}

func (list ListMem) Len() int { return len(list) }

func (list ListMem) Less(i, j int) bool {
	if list[i].role > list[j].role {
		return true
	} else if list[i].role == list[j].role {
		if list[i].id >= list[j].id {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (list ListMem) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list *ListMem) append(mem *Member) {
	*list = append(*list, mem)
}

func newListMem() (list ListMem) {
	return make([]*Member, 0)
}

func (list *ListMem) AddNew(id int, role Role, lan Language) {

	newMem := &Member{id, role, lan}

	list.append(newMem)
	log.Printf("Num of Members: %d", len(*list))
}

func (list *ListMem) AddMem(id int) {
	list.AddNew(id, Default, Vietnamese)
}

func (list *ListMem) delMem(id int) {
	temp := *list
	*list = func(temp ListMem, id int) ListMem {
		for i, v := range temp {
			if v.id == id {
				temp[i] = temp[len(temp)-1]
			}
		}
		return temp[:len(temp)-1]
	}(temp, id)
}

func (list *ListMem) String() string {
	sort.Sort(ListMem(*list))
	var content strings.Builder
	for i, v := range *list {
		content.WriteString(fmt.Sprintf("\n%d. %s", i+1, v.String()))
	}
	return content.String()
}

func (list ListMem) isInList(id int) (bool, *Member) {
	for _, v := range list {
		if v.id == id {
			return true, v
		}
	}

	return false, nil
}
