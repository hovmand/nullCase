<?php

$s = trim(fread(STDIN, 1024));

//echo $s; exit;

$data = json_decode($s, true);

$return = array(
	"Units" => array(),
);

$direction = array("N", "E", "S", "W");

foreach ($data["Units"] as $from) {
	$return["Units"][] = array("X" => $from["X"], "Y" => $from["Y"], "D" => $direction[rand(0,3)]);
}

echo json_encode($return);

//fputs(STDOUT, "PHP KODE!!!!!" . $s);

?>
