<?php

$time = microtime();
$time = explode(' ', $time);
$time = $time[1] + $time[0];
$start = $time;

$data = file_get_contents('data.txt');
file_put_contents('data-php.txt', $data);

$time = microtime();
$time = explode(' ', $time);
$time = $time[1] + $time[0];
$finish = $time;
$total_time = ($finish - $start);
echo 'Page generated in '.round(($total_time*1000), 4).' ms.';
