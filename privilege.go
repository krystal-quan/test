package main

func (list *listMem) isRole(id *int, role Role) bool {
	for _, v := range list {
		if v.id == id && v.role >= role {
			return true
		}
	}
	return false
}

func (list *listMem) isLead(id *int) bool {
	return isRole(Lead)
}

func (list *listMem) isAdmin(id *int) bool {
	return isRole(CoLead)
}

func (list *listMem) isMem(id *int) bool {
	return isRole(Member)
}
