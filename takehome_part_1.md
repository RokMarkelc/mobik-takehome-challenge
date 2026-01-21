
## 2.3 Scheduled run
Make your script run hourly. Write down what you used and how you're checking for possible errors and script failures.


Good luck!

### 1. task
- First I checked which users were logged in and their process ID with `who -u`:
```
candidate-d@box-three:~$ who -u
candidate-d sshd pts/0   2026-01-21 17:25   .       2204575 (146.212.170.247)
candidate-d pts/2        2026-01-21 14:39 02:51     1238542
```

- Then I tried to kill the process that was not mine with `sudo kill 1238542`. Did not work since I did not have root privileges.

- Then I tried to close all of my user's sessions with `loginctl terminate-user candidate-d`. 

- After logging back into the machine I checked if the malicious user's session was gone, and it was:
```
candidate-d@box-three:~$  who
candidate-d sshd pts/0   2026-01-21 17:43 (146.212.170.247)
```

### 2. task
`grep -aro "FLAG{.*}" .`

- `Contrats on the find!`
- `Nice job. Good luck with the rest!`
- `message`?

### 3.1 task
```
candidate-d@box-three:~$ grep -Erl '^.{,40}&' capture/data/
capture/data/c/c/capture_0020.csv
capture/data/c/c/capture_0013.csv
capture/data/c/c/capture_0038.csv
capture/data/c/c/capture_0044.csv
capture/data/c/c/capture_0030.csv
capture/data/c/b/capture_0027.csv
capture/data/c/b/capture_0045.csv
capture/data/c/b/capture_0002.csv
capture/data/c/b/capture_0019.csv
capture/data/c/b/capture_0029.csv
capture/data/c/a/capture_0022.csv
capture/data/c/a/capture_0012.csv
capture/data/c/a/capture_0047.csv
capture/data/c/a/capture_0020.csv
capture/data/c/a/capture_0041.csv
capture/data/c/a/capture_0013.csv
capture/data/c/a/capture_0007.csv
capture/data/b/c/capture_0025.csv
capture/data/b/c/capture_0012.csv
capture/data/b/c/capture_0009.csv
capture/data/b/c/capture_0035.csv
capture/data/b/c/capture_0040.csv
capture/data/b/c/capture_0020.csv
capture/data/b/c/capture_0048.csv
capture/data/b/c/capture_0011.csv
capture/data/b/c/capture_0034.csv
capture/data/b/c/capture_0007.csv
capture/data/b/b/capture_0014.csv
capture/data/b/b/capture_0013.csv
capture/data/b/b/capture_0031.csv
capture/data/b/b/capture_0024.csv
capture/data/b/b/capture_0050.csv
capture/data/b/b/capture_0003.csv
capture/data/b/b/capture_0037.csv
capture/data/b/d/capture_0022.csv
capture/data/b/d/capture_0016.csv
capture/data/b/d/capture_0039.csv
capture/data/b/d/capture_0015.csv
capture/data/b/d/capture_0034.csv
capture/data/b/d/capture_0044.csv
capture/data/b/a/capture_0014.csv
capture/data/b/a/capture_0050.csv
capture/data/b/a/capture_0030.csv
capture/data/b/a/capture_0007.csv
capture/data/d/c/capture_0028.csv
capture/data/d/c/capture_0014.csv
capture/data/d/b/capture_0016.csv
capture/data/d/d/capture_0026.csv
capture/data/d/d/capture_0047.csv
capture/data/d/d/capture_0013.csv
capture/data/d/d/capture_0048.csv
capture/data/d/d/capture_0024.csv
capture/data/d/a/capture_0016.csv
capture/data/d/a/capture_0010.csv
capture/data/d/a/capture_0043.csv
capture/data/d/a/capture_0024.csv
capture/data/d/a/capture_0042.csv
capture/data/d/a/capture_0015.csv
capture/data/d/a/capture_0005.csv
capture/data/a/c/capture_0001.csv
capture/data/a/c/capture_0015.csv
capture/data/a/c/capture_0034.csv
capture/data/a/b/capture_0022.csv
capture/data/a/b/capture_0012.csv
capture/data/a/b/capture_0009.csv
capture/data/a/b/capture_0001.csv
capture/data/a/b/capture_0045.csv
capture/data/a/d/capture_0027.csv
capture/data/a/d/capture_0024.csv
capture/data/a/d/capture_0004.csv
capture/data/a/d/capture_0050.csv
capture/data/a/a/capture_0040.csv
capture/data/a/a/capture_0048.csv
capture/data/a/a/capture_0004.csv
capture/data/a/a/capture_0050.csv
capture/data/a/a/capture_0006.csv
capture/data/a/a/capture_0007.csv

```

### 3.2 task
```
#!/usr/bin/env bash

# Usage: ./remove_corrupted_lines.sh

echo "$(date)" 
echo "Starting the script"

cd ..

# Checking if there are any corrupted lines
lines_to_remove=$(grep -Erh '^.{,40}$' data/ | wc -l)

if [[ $lines_to_remove -gt 0 ]]
then
  echo "Storing corrupted lines to corrupted.csv."
  grep -Erh '^.{,40}$' data/ >> corrupted.csv

  echo "Removing $lines_to_remove corrupted lines from data."
  find data -type f | xargs sed -i -r '/^.{,40}$/d'
else
  echo "There are no corrupted lines in data. Skipping removal."
fi

echo "Done."

```

### 3.3 task
I used a cron job to run the script hourly. First I edited crontab file using `crontab -e`. I added this line:
```
0 * * * * cd /home/candidate-d/capture/scripts && ./remove_corrupted_lines.sh >> /home/candidate-d/capture/logs/remove_corrupted_lines.log
```
This will make the `remove_corrupted_lines.sh` script run every hour and write the output of the script to `remove_corrupted_lines.log` where we will be able to see any errors the scripts might produce. I currently save the log file to `~/capture` for ease of use, the standard place would be somewhere in `/var/log`.

As for being alerted when this script would fail, a relatively simple way I know of is the [Healthchecks](https://healthchecks.io/docs/) service. It listens for http requests and then alerts (on Slack, mail,...) if it does not receive them at the usual time. It is a lightweight service that can be self-hosted. To make the cron job ping this service I would update the `crontab` line to:
```
0 * * * * cd /home/candidate-d/capture/scripts && ./remove_corrupted_lines.sh >> /home/candidate-d/capture/logs/remove_corrupted_lines.log && curl -fsS -m 10 --retry 5 <ping-url-generated-by-healthchecks-service>
```

If the script failed the `curl` command would not be run, Healthcheck would not get the ping it expects and the alert would go off.
