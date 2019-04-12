package simvolio

import (
	"fmt"

	"github.com/AplaProject/apla-lang/compiler"
	"github.com/AplaProject/apla-lang/runtime"
)

const (
	DEFAULT_GAS_LIMIT = 2000000000

	errCntExists    = `Contract %s has already been defined`
	errCntNotExists = `Contract %s doesn't exists`
	errGlobVar      = `%s is not a global variable`
	errGlobType     = `%s has unsupported type`
)

const (
	Void  = iota
	Int   // int
	Bool  // bool
	Str   // str
	Arr   // arr
	Map   // map
	Float // float
	Money // money
)

type Custom struct {
	Env map[string]uint32
}

type VMSettings struct {
	Custom   *Custom
	GasLimit int64
}

// RTCustom is a structure for runtime customizing
type RTCustom struct {
	Env map[string]interface{}
}

// VM is a virtual machine structure
type VM struct {
	Contracts []*runtime.Contract
	NameSpace map[string]uint32 // common namespace
	Settings  VMSettings
	Custom    *runtime.Custom
}

// NewVM creates a new virtual machine
func NewVM(settings VMSettings) *VM {
	if settings.GasLimit == 0 {
		settings.GasLimit = DEFAULT_GAS_LIMIT
	}
	env := make(map[string]runtime.EnvItem)
	var ind int
	for key, val := range settings.Custom.Env {
		env[key] = runtime.EnvItem{
			Index: ind,
			Type:  val,
		}
		ind++
	}
	return &VM{
		Contracts: make([]*runtime.Contract, 0, 1000),
		NameSpace: make(map[string]uint32),
		Settings:  settings,
		Custom: &runtime.Custom{
			Env: env,
		},
	}
}

// Compile compiles the contract and returns its structure
func (vm *VM) Compile(input string) (cnt *runtime.Contract, err error) {
	return compiler.Compile(input, &vm.NameSpace, &vm.Contracts, vm.Custom)
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
func (vm *VM) Run(cnt *runtime.Contract, Custom *RTCustom) (string, int64, error) {
	rt := runtime.NewRuntime(&vm.Contracts)
	env := make([]runtime.EnvVal, len(vm.Custom.Env))
	for key, val := range Custom.Env {
		var (
			ok      bool
			envItem runtime.EnvItem
			vEnv    int64
		)
		if envItem, ok = vm.Custom.Env[key]; !ok {
			return ``, 0, fmt.Errorf(errGlobVar, key)
		}
		switch v := val.(type) {
		case int64:
			vEnv = v
		case int:
			vEnv = int64(v)
		case string:
			rt.Strings = append(rt.Strings, v)
			vEnv = int64(len(rt.Strings) - 1)
		default:
			return ``, 0, fmt.Errorf(errGlobType, key)
		}
		env[envItem.Index] = runtime.EnvVal{
			Value: vEnv,
			Init:  true,
		}
	}
	rt.Custom = &runtime.RTCustom{
		Env: env,
	}
	return rt.Run(cnt.Code, nil, vm.Settings.GasLimit)
}

// RunByName executes the contract
func (vm *VM) RunByName(name string, Custom *RTCustom) (string, int64, error) {
	cnt := vm.GetContract(name)
	if cnt == nil {
		return ``, 0, fmt.Errorf(errCntNotExists, name)
	}
	return vm.Run(cnt, Custom)
}
