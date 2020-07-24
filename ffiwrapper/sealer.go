package ffiwrapper

import (
	"context"
	"io"

	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/filecoin-project/specs-storage/storage"

	"github.com/filecoin-project/sector-storage/storiface"
)

var _ Storage = &Sealer{}

func New(sectors SectorProvider, cfg *Config) (*Sealer, error) {
	panic("")
}

func (sb *Sealer) NewSector(ctx context.Context, sector abi.SectorID) error {
	panic("")
}

func (sb *Sealer) AddPiece(ctx context.Context, sector abi.SectorID, existingPieceSizes []abi.UnpaddedPieceSize, pieceSize abi.UnpaddedPieceSize, file storage.Data) (abi.PieceInfo, error) {
	panic("")
}

func (sb *Sealer) pieceCid(in []byte) (cid.Cid, error) {
	panic("")
}

func (sb *Sealer) UnsealPiece(ctx context.Context, sector abi.SectorID, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize, randomness abi.SealRandomness, commd cid.Cid) error {
	panic("")
}

func (sb *Sealer) ReadPiece(ctx context.Context, writer io.Writer, sector abi.SectorID, offset storiface.UnpaddedByteIndex, size abi.UnpaddedPieceSize) error {
	panic("")
}

func (sb *Sealer) SealPreCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, pieces []abi.PieceInfo) (out storage.PreCommit1Out, err error) {
	panic("")
}

func (sb *Sealer) SealPreCommit2(ctx context.Context, sector abi.SectorID, phase1Out storage.PreCommit1Out) (storage.SectorCids, error) {
	panic("")
}

func (sb *Sealer) SealCommit1(ctx context.Context, sector abi.SectorID, ticket abi.SealRandomness, seed abi.InteractiveSealRandomness, pieces []abi.PieceInfo, cids storage.SectorCids) (storage.Commit1Out, error) {
	panic("")
}

func (sb *Sealer) SealCommit2(ctx context.Context, sector abi.SectorID, phase1Out storage.Commit1Out) (storage.Proof, error) {
	panic("")
}

func (sb *Sealer) FinalizeSector(ctx context.Context, sector abi.SectorID, keepUnsealed []storage.Range) error {
	panic("")
}

func (sb *Sealer) ReleaseUnsealed(ctx context.Context, sector abi.SectorID, safeToFree []storage.Range) error {
	panic("")
}

func (sb *Sealer) Remove(ctx context.Context, sector abi.SectorID) error {
	panic("")
}

func GeneratePieceCIDFromFile(proofType abi.RegisteredSealProof, piece io.Reader, pieceSize abi.UnpaddedPieceSize) (cid.Cid, error) {
	panic("")
}

func GenerateUnsealedCID(proofType abi.RegisteredSealProof, pieces []abi.PieceInfo) (cid.Cid, error) {
	panic("")
}
