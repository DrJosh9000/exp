package algebra

var (
	// ℤ is a ring.
	_ Ring[int] = Integer{}
	// ℝ is a field.
	_ Field[float64] = Real{}
	// ℂ is a field.
	_ Field[complex128] = Cmplx{}
)

// Integer implements ℤ using int.
type Integer = BuiltinRing[int]

// Real implements ℝ using float64.
type Real = BuiltinField[float64]

// Cmplx implements ℂ using complex128 (not with the Complex type).
type Cmplx = BuiltinField[complex128]

// Fieldable supports built-in arithmetic that can support a field.
type Fieldable interface {
	~float32 | ~float64 | ~complex64 | ~complex128
}

// Fieldable supports built-in arithmetic that can support a ring.
type Ringable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | Fieldable
}

// BuiltinField implements a ring over T using built-in arithmetic.
type BuiltinRing[T Ringable] struct{}

// Add returns x+y.
func (BuiltinRing[T]) Add(x, y T) T { return x + y }

// Neg returns -x.
func (BuiltinRing[T]) Neg(x T) T { return -x }

// Zero returns 0.
func (BuiltinRing[T]) Zero() T { return 0 }

// Mul returns x*y.
func (BuiltinRing[T]) Mul(x, y T) T { return x * y }

// Identity returns 1.
func (BuiltinRing[T]) Identity() T { return 1 }

// BuiltinField implements a field over T using built-in arithmetic.
type BuiltinField[T Fieldable] BuiltinRing[T]

// Add returns x+y.
func (BuiltinField[T]) Add(x, y T) T { return x + y }

// Neg returns -x.
func (BuiltinField[T]) Neg(x T) T { return -x }

// Zero returns 0.
func (BuiltinField[T]) Zero() T { return 0 }

// Mul returns x*y.
func (BuiltinField[T]) Mul(x, y T) T { return x * y }

// Identity returns 1.
func (BuiltinField[T]) Identity() T { return 1 }

// Inv returns 1/x. This panics if x is zero.
func (BuiltinField[T]) Inv(x T) T { return 1 / x }