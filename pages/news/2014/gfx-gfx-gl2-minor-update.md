# gfx, gfx/gl2: minor update
<p class="date">October 18, 2014</p>

Recently two bugs were discovered that caused some matrices to be calculated when they shouldn't have been. We've fixed the issue in each respective package and have released the fix in a new minor version.

# gfx: version 1.0.1

There was a small bug that effectively caused any `Transform` who had a `nil` parent transform to have it's *transformation matrix* recalculated every time it was requested. Although the calculation is not very expensive -- it can be noticable in some cases (hence why the calculated transformation matrix is cached in the first place).

See the actual issue for this bug [here](https://github.com/azul3d/gfx/issues/16).

# gfx/gl2: version 2.0.1

In a similar nature to the above -- the gfx/gl2 package performs the calculation of what is called the *model view projection (MVP) matrix*. This is a simple three 4x4 matrix multiplications -- which again isn't very expensive. But to get the best performance we do have a caching system to ensure that we don't calculate that matrix each time an object is drawn. We discovered a small bug that rendered the cache ineffective and have fixed it.

See the actual issue for this bug [here](https://github.com/azul3d/gfx-gl2/issues/22).

# Updating

To update all you need to do is run the following (*import paths are the same*):

```
go get -u azul3d.org/gfx.v1
go get -u azul3d.org/gfx/gl2.v2
```

# Conclusion

Users should expect to see a small performance gain when updating. If your application makes heavy use of *transforms* then you should see a significant performance increase.
