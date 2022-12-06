package main

import "github.com/ruijzhan/demo/http/framework"

func SubjectAddController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectAddController")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectListController")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectDelController")
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectUpdateController")
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectGetController")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.SetStatus(200).Json("ok, SubjectNameController")
	return nil
}
