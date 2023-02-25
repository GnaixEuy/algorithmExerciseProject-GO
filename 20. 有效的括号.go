package main

import "container/list"

func isValid(s string) bool {
	l := list.New()
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			l.PushBack(ch)
		} else {
			if l.Back() == nil {
				return false
			}
			if ch == ')' && l.Back().Value == '(' {
				top := l.Back()
				l.Remove(top)
			} else if ch == '}' && l.Back().Value == '{' {
				top := l.Back()
				l.Remove(top)
			} else if ch == ']' && l.Back().Value == '[' {
				top := l.Back()
				l.Remove(top)
			} else {
				return false
			}
		}
	}
	return l.Len() == 0
}
