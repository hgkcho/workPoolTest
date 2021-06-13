package main

const JobCount = 100

func main() {
	d := newDispatchar()
	d.Start()
	for _, v := range createJobs(JobCount) {
		d.Add(v)
	}
	d.Wait()
}
