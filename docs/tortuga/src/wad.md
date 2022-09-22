# WAD
`WAD` files are simply a mapping of files that contain data used to build a scenel.

# Tiles
`tiles` is a mapping of which character will represent which block of pixels in our background.
These chars are used to easily build backgrounds on a tile layer.

For example, 
A `.map` file may contain a set of chars that will get parsed and rendered as tiles in the background.


# Layers
A `.map` file represents where each block of pixels will exist.

These tiles can be loaded into tile memory which may be rerendered less than other things on the screen.

# Sprites
Sprites can be loaded from `.spr` files.
Sprites are more likely to be rerendered more often than tiles.

## Sprite Encoding
A pixel is represented by a hexadecimal char.  The hexidecimal character is an index on the palette.  So, there can only be 16 colors in the palette.
