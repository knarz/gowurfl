# gowurfl

gowurfl is a wrapper around the C API exposed by libwurfl distributed/sold
by (scienta mobile)[https://www.scientiamobile.com]. An outdated documentation
can be found (here)[https://docs.scientiamobile.com/documentation/infuze/infuze-c++-api-user-guide].

The code is based on the documentation found in the `wurfl/wurfl.h` header file.

To use this library you need to have libwurfl installed in such a way that
the cgo compiler can find it. The only tested version at this point is `1.7.1.0`.
