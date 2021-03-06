package bitflag

// Thanks to prim, he is the one who made this bitflag thing in the first place.
// I just reused his code and made an actual manager.
// https://github.com/Prim69

type FlagManager struct {
	flags uint64
}

func NewManager() *FlagManager { return &FlagManager{flags: 0} }

// HasFlag returns whether the manager has a specified bitflag.
func (m *FlagManager) HasFlag(flag uint64) bool {
	return m.flags&flag > 0
}

// HasAnyFlag returns whether the manager has any of the specified bitflags.
func (m *FlagManager) HasAnyFlag(flags ...uint64) bool {
	for _, flag := range flags {
		if m.HasFlag(flag) {
			return true
		}
	}
	return false
}

// HasAllFlags returns true if the manager has all the flags specified.
func (m *FlagManager) HasAllFlags(flags ...uint64) bool {
	for _, flag := range flags {
		if !m.HasFlag(flag) {
			return false
		}
	}
	return true
}

// setFlag sets a bit flag for the manager, or unsets if the manager already has the flag.
func (m *FlagManager) setFlag(flag uint64) {
	m.flags ^= flag
}

// AddFlag will add a bit flag based on the value of set.
func (m *FlagManager) AddFlag(flag uint64) {
	if !m.HasFlag(flag) {
		m.setFlag(flag)
	}
}

// RemoveFlag will remove a bit flag based on the value of set.
func (m *FlagManager) RemoveFlag(flag uint64) {
	if m.HasFlag(flag) {
		m.setFlag(flag)
	}
}

// Flags will return the bit flags of the manager.
func (m *FlagManager) Flags() uint64 { return m.flags }
