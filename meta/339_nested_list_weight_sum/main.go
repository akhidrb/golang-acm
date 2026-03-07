package main

type NestedInteger struct {
	value *int
	list  []*NestedInteger
}

func NewInteger(val int) *NestedInteger {
	return &NestedInteger{value: &val}
}

func NewList(list []*NestedInteger) *NestedInteger {
	return &NestedInteger{list: list}
}

func (n *NestedInteger) IsInteger() bool {
	return n.value != nil
}

func (n *NestedInteger) GetInteger() int {
	if n.value == nil {
		return 0
	}
	return *n.value
}

func (n *NestedInteger) GetList() []*NestedInteger {
	return n.list
}

func depthSum(nestedList []*NestedInteger) int {
	return dfs(nestedList, 1)
}

func dfs(nestedList []*NestedInteger, depth int) int {
	sum := 0

	for _, value := range nestedList {
		if value.IsInteger() {
			sum += value.GetInteger() * depth
		} else {
			sum += dfs(value.GetList(), depth+1)
		}
	}
	return sum
}

func depthSum2(nestedList []*NestedInteger) int {
	queue := nestedList
	sum := 0
	depth := 1

	for len(queue) > 0 {
		next := make([]*NestedInteger, 0)

		for _, value := range queue {
			if value.IsInteger() {
				sum += value.GetInteger() * depth
			} else {
				next = append(next, value.GetList()...)
			}
		}

		queue = next
		depth++

	}
	return sum
}
