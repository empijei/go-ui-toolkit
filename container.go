package goUIToolKit

type Container interface {
	Component
	AddComponent(Component)
}
