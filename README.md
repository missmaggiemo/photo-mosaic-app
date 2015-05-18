# Simple Photo Mosaic App

Nothing fancy. Run `source main.sh` to set up your `bee` path and `cd` you into the project directory.

### The Challenge

*Preamble*

A photographic mosaic, or a photo-mosaic is a picture (usually a photograph) that has been divided into (usually equal sized) rectangular sections, each of which is replaced with another picture (called a tile picture). If we view it from far away or if you squint at it, then the original picture can be seen. If we look closer though, we will see that the picture is in fact made up of many hundreds or thousands of smaller tile pictures.

*Goals of the challenge*

Your mission, should you accept it, is to write a photo-mosaic generating program that:

* Allows the user to select a target picture, which is the picture that will be made into a photo-mosaic
* Allows the user to select a directory containing a set of tile pictures
* Generates a photo-mosaic of the target picture using the tile pictures
* Bonus goals (optional, not part of the challenge)

Create a web application that generates the photo-mosaic that:

* Allows the user to log in (can be your own database or log in through a third party like GitHub or Twitter or Facebook, through OAuth2). (~Note: if you are authenticating the user through OAuth2brary).
* Allows the user to connect to one or more photo-sharing sites like Instagram or Flickr or Facebook Photos (or any photo-sharing site of your choice) to get tile pictures. Your user doesn't necessarily need to log in, you can use the image search APIs to get the tile pictures
* Allows the user to use a search filter (for e.g. use only pictures with cats in it) to filter out a set of tile pictures
* Allows the user to save the photo-mosaic, either on the site or upload it back to the photo-sharing site

*Requirements of the challenge*

* Use the latest version of Go i.e. version 1.4.2
* Individual tile pictures must be clearly visible when magnified, though it is expected to be smaller.
* You need to write test cases for the main flow. Do submit your test cases.
* Do organize your code.