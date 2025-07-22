<?php

for ($i = 0; $i < 20; $i++) {
    $vetor[$i] = (int) trim(fgets(STDIN));
}

// $vetor = array_reverse($vetor);

for ($i = 0; $i < 20; $i++) {
    $vetor[$i] = $vetor[19 - $i];
    echo "N[$i] = {$vetor[$i]}\n";
}

?>
