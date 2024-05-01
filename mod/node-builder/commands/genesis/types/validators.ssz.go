// Code generated by fastssz. DO NOT EDIT.
// Hash: 6bba140b7126ff0f0d56770a0e8c7e64ff59f3e4882ad82fbc8e242d96859b11
// Version: 0.1.3
package types

import (
	"github.com/berachain/beacon-kit/mod/primitives/pkg/consensus"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the ValidatorsMarshaling object
func (v *ValidatorsMarshaling) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(v)
}

// MarshalSSZTo ssz marshals the ValidatorsMarshaling object to a target array
func (v *ValidatorsMarshaling) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(4)

	// Offset (0) 'Validators'
	dst = ssz.WriteOffset(dst, offset)

	// Field (0) 'Validators'
	if size := len(v.Validators); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("ValidatorsMarshaling.Validators", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(v.Validators); ii++ {
		if dst, err = v.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	return
}

// UnmarshalSSZ ssz unmarshals the ValidatorsMarshaling object
func (v *ValidatorsMarshaling) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 4 {
		return ssz.ErrSize
	}

	tail := buf
	var o0 uint64

	// Offset (0) 'Validators'
	if o0 = ssz.ReadOffset(buf[0:4]); o0 > size {
		return ssz.ErrOffset
	}

	if o0 < 4 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (0) 'Validators'
	{
		buf = tail[o0:]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		v.Validators = make([]*consensus.Validator, num)
		for ii := 0; ii < num; ii++ {
			if v.Validators[ii] == nil {
				v.Validators[ii] = new(consensus.Validator)
			}
			if err = v.Validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the ValidatorsMarshaling object
func (v *ValidatorsMarshaling) SizeSSZ() (size int) {
	size = 4

	// Field (0) 'Validators'
	size += len(v.Validators) * 121

	return
}

// HashTreeRoot ssz hashes the ValidatorsMarshaling object
func (v *ValidatorsMarshaling) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(v)
}

// HashTreeRootWith ssz hashes the ValidatorsMarshaling object with a hasher
func (v *ValidatorsMarshaling) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(v.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range v.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the ValidatorsMarshaling object
func (v *ValidatorsMarshaling) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(v)
}
