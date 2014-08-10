gomasticate
=========
[![Build Status](https://magnum.travis-ci.com/CapillarySoftware/gomasticate.svg?token=48o3zC2UVnhZLFcYLG2C&branch=master)](https://magnum.travis-ci.com/CapillarySoftware/gomasticate)

Masticate... because it feels right.

<h3>install</h3>
<pre>
<code>
//nanomsg
http://nanomsg.org/download.html
./configure && make && make check && sudo make install
go get github.com/tools/godep
godep restore
go get github.com/op/go-nanomsg
go get github.com/CapillarySoftware/gomasticate
</code>
</pre>

<h3>Install statically linked version</h3>
<pre>
<code>
go install --ldflags '-extldflags "-static"'  github.com/CapillarySoftware/gomasticate
</code>
</pre>
