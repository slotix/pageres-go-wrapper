package sshot

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//Parameters struct serves for passing parameters for pageres
//Read more information about parameters at https://github.com/sindresorhus/pageres
type Parameters struct {
	Command   string //pageres
	Sizes     string //Use a <width>x<height> notation or a keyword.
	Crop      string //Crop to the set height. add --crop to args
	Scale     string //Scale webpage n times.
	Timeout   string //Number of seconds after which PhantomJS aborts the request.
	Filename  string //Define a customized filename. For example <%= date %> - <%= url %>-<%= size %><%= crop %>.
	UserAgent string //Custom user agent.
}

//GetShots starts pageres process itself and returns URL if Pageres Couldn't load url.
//Empty string is returned overwise
func GetShots(urls []string, params Parameters) {
	//sometimes pageres returns an error "Couldn't load url" for some url
	//in this case all other urls is not processed
	//In order to process the rest except bad urls
	//a badURL need to be removed from a list and start again
	success := false
	for success == false {
		badURL := runShotClient(urls, params)
		if badURL == "" {
			success = true
		} else {
			urls = deleteStringFromSlice(badURL, urls)
		}
	}
}

func runShotClient(urls []string, params Parameters) string {
	var args []string
	command := params.Command
	//pageres http://google.com https://dbconvert.com '480x320' --crop --scale 0.9 --filename 'media/shots/<%= url %>'
	args = append(
		urls,
		params.Sizes,
		params.Crop,
		params.Scale,
		params.Timeout,
		params.Filename,
	)
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
		//if process of even one of urls returns "Couldn't load url"
		//pageres stops processing other ones.
		if strings.Contains(stderr.String(), "Couldn't load url") {
			return getBadURLFromErr(stderr.String())
		}
	}
	return ""
}

//getBadURLFromErr gets url from stderr.String()
//for example exit status 1: Couldn't load url: http://baguioautoforums.com/index.php
func getBadURLFromErr(strErr string) string {
	re := regexp.MustCompile("https?.+")
	found := re.FindString(strErr)
	return found
}

//DeleteZeroLengthFiles  deletes zero length files from a specified dir
func DeleteZeroLengthFiles(dir string) {
	//find /foo/bar -size 0c -delete
	command := "find"
	var args []string
	args = append(args, dir, "-size", "0c", "-delete")
	out, err := exec.Command(command, args...).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Printf(string(out))
	}
}

//deleteStringFromSlice Delete string from slice
func deleteStringFromSlice(d string, s []string) []string {
	var r []string
	for _, str := range s {
		if str != d {
			r = append(r, strings.TrimSpace(str))
		}
	}
	return r
}
