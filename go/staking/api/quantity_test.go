package api

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"
)

func fromInt(n int) *Quantity {
	q := NewQuantity()
	q.inner.SetInt64(int64(n))
	return q
}

func (q *Quantity) eqInt(n int) bool {
	nq := fromInt(n)
	return q.Cmp(nq) == 0
}

func TestQuantityCtors(t *testing.T) {
	require := require.New(t)

	q := NewQuantity()
	require.NotNil(q, "NewQuantity")
	require.True(q.eqInt(0), "New value")

	q = fromInt(23)
	nq := q.Clone()
	_ = q.FromBigInt(big.NewInt(666))
	require.True(nq.eqInt(23), "Clone value")
}

func TestFromBigInt(t *testing.T) {
	require := require.New(t)

	var q Quantity
	err := q.FromBigInt(nil)
	require.Equal(ErrInvalidArgument, err, "FromBigInt(nil)")

	err = q.FromBigInt(big.NewInt(-1))
	require.Equal(ErrInvalidArgument, err, "FromBigInt(-1)")

	err = q.FromBigInt(big.NewInt(23))
	require.NoError(err, "FromBigInt(23)")
	require.True(q.eqInt(23), "FromBigInt(23) value")
}

func TestQuantityBinaryRoundTrip(t *testing.T) {
	const expected int = 0xdeadbeef

	require := require.New(t)

	q := fromInt(expected)
	b, err := q.MarshalBinary()
	require.NoError(err, "MarshalBinary")

	var nq Quantity
	err = nq.UnmarshalBinary(b)
	require.NoError(err, "UnmarshalBinary")

	require.Zero(q.Cmp(&nq), "Round trip matches")
}

func TestQuantityAdd(t *testing.T) {
	require := require.New(t)

	q := fromInt(100)

	err := q.Add(nil)
	require.Equal(ErrInvalidArgument, err, "Add(nil)")

	err = q.Add(fromInt(-1))
	require.Equal(ErrInvalidArgument, err, "Add(-1)")

	err = q.Add(fromInt(200))
	require.NoError(err, "Add")
	require.True(q.eqInt(300), "Add(200) value")
}

func TestQuantitySub(t *testing.T) {
	require := require.New(t)

	q := fromInt(100)

	err := q.Sub(nil)
	require.Equal(ErrInvalidArgument, err, "Sub(nil)")

	err = q.Sub(fromInt(-1))
	require.Equal(ErrInvalidArgument, err, "Sub(-1)")

	err = q.Sub(fromInt(200))
	require.Equal(ErrInsufficientBalance, err, "Sub(200)")

	err = q.Sub(fromInt(23))
	require.NoError(err, "Sub")
	require.True(q.eqInt(77), "Sub(23) value")
}

func TestQuantitySubUpTo(t *testing.T) {
	require := require.New(t)

	q := fromInt(100)

	_, err := q.SubUpTo(nil)
	require.Equal(ErrInvalidArgument, err, "SubUpTo(nil)")

	_, err = q.SubUpTo(fromInt(-1))
	require.Equal(ErrInvalidArgument, err, "SubUpTo(-1)")

	n, err := q.SubUpTo(fromInt(23))
	require.NoError(err, "SubUpTo")
	require.True(q.eqInt(77), "SubUpTo(23) value")
	require.True(n.eqInt(23), "SubUpTo(23) subtracted")

	n, err = q.SubUpTo(fromInt(9000))
	require.NoError(err, "SubUpTo(9000)")
	require.True(q.eqInt(0), "SubUpTo(9000) value")
	require.True(n.eqInt(77), "SubUpTo(9000) subtracted")
}

func TestQuantityCmp(t *testing.T) {
	require := require.New(t)

	q := fromInt(100)

	require.Equal(-1, q.Cmp(fromInt(9001)), "q.Cmp(9001)")
	require.Equal(0, q.Cmp(fromInt(100)), "q.Cmp(100)")
	require.Equal(1, q.Cmp(fromInt(42)), "q.Cmp(42)")

	require.False(q.IsZero(), "q.IsZero()")
	require.True(NewQuantity().IsZero(), "NewQuantity().IsZero()")
}

func TestQuantityString(t *testing.T) {
	require := require.New(t)

	require.Equal("-500", fromInt(-500).String(), "Invalid returns raw inner")
	require.Equal("123456", fromInt(123456).String(), "Positive integer")
}

func TestMove(t *testing.T) {
	require := require.New(t)

	err := Move(nil, fromInt(100), fromInt(25))
	require.Equal(err, ErrInvalidAccount, "Move(nil, 100, 25)")
	err = Move(fromInt(50), nil, fromInt(25))
	require.Equal(err, ErrInvalidAccount, "Move(50, nil, 25)")
	err = Move(fromInt(50), fromInt(100), nil)
	require.Equal(err, ErrInvalidArgument, "Move(50, 100, nil)")

	dst, src := fromInt(100), fromInt(300)
	err = Move(dst, src, fromInt(9000))
	require.Equal(err, ErrInsufficientBalance, "Move(100, 300, 9000)")
	require.True(dst.eqInt(100) && src.eqInt(300), "Move(fail) - dst/src unchanged")

	err = Move(dst, src, fromInt(75))
	require.NoError(err, "Move")
	require.True(dst.eqInt(175), "Move - dst value")
	require.True(src.eqInt(225), "Move - src value")
}

func TestMoveUpTo(t *testing.T) {
	require := require.New(t)

	_, err := MoveUpTo(nil, fromInt(100), fromInt(25))
	require.Equal(err, ErrInvalidAccount, "MoveUpTo(nil, 100, 25)")
	_, err = MoveUpTo(fromInt(50), nil, fromInt(25))
	require.Equal(err, ErrInvalidAccount, "MoveUpTo(50, nil, 25)")
	_, err = MoveUpTo(fromInt(50), fromInt(100), nil)
	require.Equal(err, ErrInvalidArgument, "MoveUpTo(50, 100, nil)")

	dst, src := fromInt(100), fromInt(300)
	moved, err := MoveUpTo(dst, src, fromInt(75))
	require.NoError(err, "MoveUpTo")
	require.True(dst.eqInt(175), "MoveUpTo - dst value")
	require.True(src.eqInt(225), "MoveUpTo - src value")
	require.True(moved.eqInt(75), "MoveUpTo - moved")

	moved, err = MoveUpTo(dst, src, fromInt(90000))
	require.NoError(err, "MoveUpTo, oversized")
	require.True(dst.eqInt(400), "MoveUpTo, oversized - dst value")
	require.True(src.eqInt(0), "MoveUpTo, oversized - src value")
	require.True(moved.eqInt(225), "MoveUpTo, oversized - moved")
}