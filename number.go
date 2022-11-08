package deagon

// Name returns a name from an index. Only low 25 bits from index are used.
func Name(formatter Formatter, index int) string {
  return getName(index, formatter)
}
