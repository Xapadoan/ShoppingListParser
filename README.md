# SHoPing LiSt PaRSeR

Reads a single string and parses it as quantity and ingredient, e.g 200g sugar becomes:
 - 200 (quantity - number)
 - g (unit - enum)
 - ingredient (name - string)

### Recognized Units

This program will recognize the following symbols as units:
 - Nothing or "u" -> Units
 - "g" -> Gram
 - "l" or "L" -> Liter
 - "ml" or "mL" -> Milliliter
 - "cl" or "cL" -> Centiliter