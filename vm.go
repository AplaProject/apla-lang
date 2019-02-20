package simvolio

import (
	"fmt"

	"github.com/AplaProject/apla-lang/compiler"
	"github.com/AplaProject/apla-lang/runtime"
)

const (
	errCntExists    = `Contract %s has already been defined`
	errCntNotExists = `Contract %s doesn't exists`
)

// VM is a virtual machine structure
type VM struct {
	Contracts []*runtime.Contract
	NameSpace map[string]uint32 // common namespace
}

// NewVM creates a new virtual machine
func NewVM() *VM {
	return &VM{
		Contracts: make([]*runtime.Contract, 0, 1000),
		NameSpace: make(map[string]uint32),
	}
}

// Compile compiles the contract and returns its structure
func (vm *VM) Compile(input string) (cnt *runtime.Contract, err error) {
	return compiler.Compile(input, &vm.NameSpace, &vm.Contracts)
}

// GetContract returns the contract by its name
func (vm *VM) GetContract(name string) *runtime.Contract {
	if ind, ok := vm.NameSpace[name]; ok {
		return vm.Contracts[ind]
	}
	return nil
}

// Link links the compiled contract to VM
func (vm *VM) Link(cnt *runtime.Contract, reload bool) error {
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
		ind = uint32(len(vm.Contracts) - 1)
	}
	vm.NameSpace[cnt.Name] = ind
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
func (vm *VM) Run(cnt *runtime.Contract) (string, int64, error) {
	rt := runtime.NewRuntime(&vm.Contracts)
	return rt.Run(cnt.Code, nil)
}

// RunByName executes the contract
func (vm *VM) RunByName(name string) (string, int64, error) {
	cnt := vm.GetContract(name)
	if cnt == nil {
		return ``, 0, fmt.Errorf(errCntNotExists, name)
	}
	return vm.Run(cnt)
}
