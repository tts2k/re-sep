# RE:SEP
An alternative article reader for the [Stanford Encyclopedia of
Philosophy](https://plato.stanford.edu). Heavily inspired by
[Foliate](https://johnfactotum.github.io/foliate/).


## Planning
The architecture is pretty much near finalized, as in, it has reached a desired
level of modularity and abstraction for me.

The project started as a way for me to learn new platforms and frameworks, and
it will continue to be my learning toy since I can basically replace any parts of
it with something else as long as the message passing methods are available in
that language and can be properly implemented.

Learning is a non-stop process and so are new javascript framework releases.

### Architecture
[![architecture](docs/architecture.png)](docs/architecture.png)
Click on the image for full view

### Pipeline
![pipeline](docs/pipeline.png)
Notes:
- [1]: The current plan is to do user config styling based on pre-defined
values. Those value will be represented in styles using custom Tailwind classes.
The Golang backend will use goquery to add those classes to the html text before
sending out
- [2]: For user config changes from front-end, beside updating user config to
the database, the style changes will be hot-applied to current page using
client-side javascript. This is an attempt to both reduce server load and
improve UX. I call this process **Layering**, inspired by the
[similar functionality](https://docs.fedoraproject.org/en-US/fedora-silverblue/getting-started/#package-layering)
of Fedora Silverblue.

### Status: Arborescent
It's a tree!
