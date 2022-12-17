<?php

// memanggil function main untuk memulai menjalankan program
main();

// function untuk menampung kode utama, agar lebih rapi
function main(): void {
    // membaca isi file dataset.json dan menampungnya ke dalam variable $file
    $file = file_get_contents("dataset.json");

    // mengubah tipe data variable $file menjadi array associative dan menampungnya di variable $dataset
    $dataset = json_decode($file, true);

	// normalisasi nilai bobot
	$dataset = weightNormalization($dataset);

	// normalisasi nilai matriks
    $dataset = matrixNormalization($dataset);

	// menghitung pembobotan matriks yang telah dinormalisasi
    $dataset = resultCalculation($dataset);

	// mengurutkan hasil berdasarkan pembobotan matriks tertinggi
    usort($dataset['alternatives'], fn($alt1, $alt2) => (int) ($alt1['result'] <=> $alt2['result']));
    $dataset['alternatives'] = array_reverse($dataset['alternatives']);

	// menampilkan hasil menggunakan std out CLI
    echo '================================' . PHP_EOL;
    echo "No.\tNama \t\tHasil" . PHP_EOL;
    echo '================================' . PHP_EOL;
    foreach ($dataset['alternatives'] as $index => $alternative) {
        echo sprintf("%2d\t%-10s \t %.2f\n", $index+1, $alternative['name'], $alternative['result']);
    }
}

// function untuk menormalisasi nilai bobot
function weightNormalization(array $dataset): array {
	// Menghitung total bobot
    $total = 0.0;
    foreach ($dataset['criteria'] as $criterion) {
        $total += (float) $criterion['weight'];
    }

	// Menghitung hasil normalisasi masing-masing bobot
    foreach ($dataset['criteria'] as $index => $criterion) {
        $dataset['criteria'][$index]['normalized_weight'] = (float) $criterion['weight'] / $total;
    }

    return $dataset;
}

// function untuk menormalisasi nilai matriks
function matrixNormalization(array $dataset): array {
	// menentukan nilai pembagi untuk masing-masing kriteria berdasarkan nilai-nilai alternatif
    foreach ($dataset['alternatives'] as $alternative) {
        foreach ($dataset['criteria'] as $index => $criterion) {
            $dataset['criteria'][$index]['divisor_value'] = divisorValue($alternative[$criterion['code']], $criterion['divisor_value'] ?? 0, $criterion['type']);
        }
    }

	// normalisasi matriks berdasarkan nilai pembaginya
    foreach ($dataset['alternatives'] as $index => $alternative) {
        foreach ($dataset['criteria'] as $criterion) {
            $dataset['alternatives'][$index]['normalized_' . $criterion['code']] = normalize($alternative[$criterion['code']], $criterion['divisor_value'], $criterion['type']);
        }
    }

    return $dataset;
}

// function untuk menghitung pembobotan matriks yang telah dinormalisasi
function resultCalculation(array $dataset): array {
    foreach ($dataset['alternatives'] as $index => $alternative) {
        $dataset['alternatives'][$index]['result'] = 0;
        foreach ($dataset['criteria'] as $criterion) {
            $dataset['alternatives'][$index]['result'] += (float) ($alternative['normalized_' . $criterion['code']] * $criterion['normalized_weight']);
        }
    }
    return $dataset;
}

// function untuk menghitung nilai normalisasi matriks
function normalize(float $matrix, float $divisor, string $type): float {
    if ($type == 'cost') {
        return (float) $divisor / $matrix;
    }

    return $matrix / $divisor;
}

// function untuk mendapatkan pembagi yang akan digunakan untuk normalisasi matriks
function divisorValue(float $value, float $initial, string $type): float {
    if ($type == 'cost' && ($value < $initial || $initial == 0)) {
        return $value;
    }

    if ($type == 'benefit' && $value > $initial) {
        return $value;
    }

	return $initial;
}
