import std.stdio;

/+ dub.sdl:
dependency "vibe-d" version="~>0.9.0"
+/
void main()
{
    import vibe.d;
    listenHTTP(":8080", (req, res) {
        res.writeBody("Hello, World: " ~ req.path);
    });
    runApplication();
}
