# Features

[] Merging logfiles for multiple jobs
- per_job_logs=0 might help combine the logs but https://github.com/axboe/fio/issues/1032 indicates this might not work with client/server
[] support extra metadata about cloud vendor volume type, iops limits, etc
[] option to only create output.json instead of logs
- Logs can be large and slow to copy from client/server setups
- sometimes the summaries are enough for some needs
[] alternate graphing (plotinum?) because i don't really want to figure out the js/d3 app
[] storing multiple runs of the same suite - e.g. sub-dirs by timestamp?
[] summarizing multiple runs of the same suite - e.g. an average of the average iops over all the runs (would only be useful for steady-state tests, not tests with phases)
   * Why? For avoiding outliers from cloud services where performance can vary.  May want to run multiple times to confirm that it wasn't a fluke
[] display/link fio benchmark details from the summary
[] ssh to server host to ensure that `fio --server` is running

[x] support fio remote client to server
[x] client/server log files end with the server name so they aren't picked up by Inventory
# Bugs
[x] Diskstats csv is empty
[] iops graphs are just constant value of 1?  is it not summing?
[] the graph x-axis always says 0-10 minutes even for runs that are much shorter and shows results at times that they could not have been


## Notes

[x] understand what the current summary behavior is.  the graphs just display the suite name
- The device file looks like it's supposed to just contain devices from one host.  Of course, because without client/server how could it be anything else?
 - it does work to run client/server against different hosts from one device file
 - but probably need to summarize/visualize differently
- But what is the purpose of the graphs displaying a summary of all devices with no way to see individual device results?
  - generates a series for each device

### Multiple runs

* currently it looks like running a suite again may not do anything unless you delete the previous results
 * actually, use the -rerun flag
 * also, adding another device to a suite that was previously run doesn't update suite.json so it seems like 
   this expects that a suite is run all at once and not incrementally
