package simvolio

import (
	"fmt"
)

const (
	errCntExists    = `Contract %s has already been defined`
	errCntNotExists = `Contract %s doesn't exists`
)

// Contract contains information about the contract
type Contract struct {
	ID   int64 // External id
	Name string
	Code []uint32
}

// VM is a virtual machine structure
type VM struct {
	Contracts []*Contract
	NameSpace map[string]uint32 // common namespace
}

// NewVM creates a new virtual machine
func NewVM() *VM {
	return &VM{
		Contracts: make([]*Contract, 0, 1000),
		NameSpace: make(map[string]uint32),
	}
}

// Compile compiles the contract and returns its structure
func (vm *VM) Compile(input string) (cnt *Contract, err error) {
	cnt = &Contract{
		Code: make([]uint32, 0, 64),
	}
	return
}

// GetContract returns the contract by its name
func (vm *VM) GetContract(name string) *Contract {
	if ind, ok := vm.NameSpace[name]; ok {
		return vm.Contracts[ind]
	}
	return nil
}

// Link links the compiled contract to VM
func (vm *VM) Link(cnt *Contract, reload bool) error {
	var (
		ind uint32
		ok  bool
	)
	if ind, ok = vm.NameSpace[cnt.Name]; ok && !reload {
		return fmt.Errorf(errCntExists, cnt.Name)
	} else if !ok && reload {
		return fmt.Errorf(errCntNotExists, cnt.Name)
	}
	if reload {
		vm.Contracts[ind] = cnt
	} else {
		vm.Contracts = append(vm.Contracts, cnt)
	}
	return nil
}

// LoadContract compiles and link the contract
func (vm *VM) LoadContract(input string, id int64) error {
	cnt, err := vm.Compile(input)
	if err != nil {
		return err
	}
	if err = vm.Link(cnt, false); err != nil {
		return err
	}
	cnt.ID = id
	return nil
}

// Run executes the contract
func (vm *VM) Run(cnt *Contract) (string, int64, error) {
	return ``, 0, nil
}

// RunByName executes the contract
func (vm *VM) RunByName(name string) (string, int64, error) {
	cnt := vm.GetContract(name)
	if cnt == nil {
		return ``, 0, fmt.Errorf(errCntNotExists, name)
	}
	return vm.Run(cnt)
}
