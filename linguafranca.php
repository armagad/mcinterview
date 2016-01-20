#!/usr/local/bin/php
<?php

array_shift($argv); // lose the program name

foreach($argv as $arg){
	foreach(str_split(strtolower($arg)) as $_){
		$counts[$_] += 1;
	}
}

array_multisort($counts, SORT_DESC, array_keys($counts));

echo json_encode($counts);
echo "\n";
