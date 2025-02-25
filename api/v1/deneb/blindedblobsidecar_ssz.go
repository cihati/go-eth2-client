// Code generated by fastssz. DO NOT EDIT.
// Hash: 53f7a78b79c00bbf864316c122774dbb8e8aae2e6c5245947a816bca4690d094
// Version: 0.1.3
package deneb

import (
	"github.com/attestantio/go-eth2-client/spec/deneb"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BlindedBlobSidecar object
func (b *BlindedBlobSidecar) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BlindedBlobSidecar object to a target array
func (b *BlindedBlobSidecar) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf

	// Field (0) 'BlockRoot'
	dst = append(dst, b.BlockRoot[:]...)

	// Field (1) 'Index'
	dst = ssz.MarshalUint64(dst, uint64(b.Index))

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (3) 'BlockParentRoot'
	dst = append(dst, b.BlockParentRoot[:]...)

	// Field (4) 'ProposerIndex'
	dst = ssz.MarshalUint64(dst, uint64(b.ProposerIndex))

	// Field (5) 'BlobRoot'
	dst = append(dst, b.BlobRoot[:]...)

	// Field (6) 'KzgCommitment'
	dst = append(dst, b.KzgCommitment[:]...)

	// Field (7) 'KzgProof'
	dst = append(dst, b.KzgProof[:]...)

	return
}

// UnmarshalSSZ ssz unmarshals the BlindedBlobSidecar object
func (b *BlindedBlobSidecar) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size != 216 {
		return ssz.ErrSize
	}

	// Field (0) 'BlockRoot'
	copy(b.BlockRoot[:], buf[0:32])

	// Field (1) 'Index'
	b.Index = deneb.BlobIndex(ssz.UnmarshallUint64(buf[32:40]))

	// Field (2) 'Slot'
	b.Slot = phase0.Slot(ssz.UnmarshallUint64(buf[40:48]))

	// Field (3) 'BlockParentRoot'
	copy(b.BlockParentRoot[:], buf[48:80])

	// Field (4) 'ProposerIndex'
	b.ProposerIndex = phase0.ValidatorIndex(ssz.UnmarshallUint64(buf[80:88]))

	// Field (5) 'BlobRoot'
	copy(b.BlobRoot[:], buf[88:120])

	// Field (6) 'KzgCommitment'
	copy(b.KzgCommitment[:], buf[120:168])

	// Field (7) 'KzgProof'
	copy(b.KzgProof[:], buf[168:216])

	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BlindedBlobSidecar object
func (b *BlindedBlobSidecar) SizeSSZ() (size int) {
	size = 216
	return
}

// HashTreeRoot ssz hashes the BlindedBlobSidecar object
func (b *BlindedBlobSidecar) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BlindedBlobSidecar object with a hasher
func (b *BlindedBlobSidecar) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'BlockRoot'
	hh.PutBytes(b.BlockRoot[:])

	// Field (1) 'Index'
	hh.PutUint64(uint64(b.Index))

	// Field (2) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (3) 'BlockParentRoot'
	hh.PutBytes(b.BlockParentRoot[:])

	// Field (4) 'ProposerIndex'
	hh.PutUint64(uint64(b.ProposerIndex))

	// Field (5) 'BlobRoot'
	hh.PutBytes(b.BlobRoot[:])

	// Field (6) 'KzgCommitment'
	hh.PutBytes(b.KzgCommitment[:])

	// Field (7) 'KzgProof'
	hh.PutBytes(b.KzgProof[:])

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BlindedBlobSidecar object
func (b *BlindedBlobSidecar) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
