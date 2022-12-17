const fs = require('fs');

// memanggil function main untuk memulai menjalankan program
main();

// function untuk menampung kode utama, agar lebih rapi
function main() {
    // membaca isi file dataset.json dan menampungnya ke dalam constant file
    const file = fs.readFileSync('dataset.json');

    // mengubah tipe data constant file menjadi object dan menampungnya di constant dataset
    const dataset = JSON.parse(file);

    // normalisasi nilai bobot
    weightNormalization(dataset);

    // normalisasi nilai matriks
    matrixNormalization(dataset);

    // menghitung pembobotan matriks yang telah dinormalisasi
    resultCalculation(dataset);

    // mengurutkan hasil berdasarkan pembobotan matriks tertinggi
    dataset?.alternatives.sort((alt1, alt2) => (alt1.result < alt2.result) ? 1 : -1);

    // menampilkan hasil menggunakan std out CLI
    console.log('================================');
    console.log("No.\tNama \t\tHasil");
    console.log('================================');
    dataset?.alternatives.forEach((alternative, index) => {
        const number = (index+1 < 10) ? ` ${index+1}` : index+1;
        console.log(`${number}\t${alternative?.name}\t\t${alternative?.result.toFixed(2)}`);
    })
}

// function untuk menormalisasi nilai bobot
function weightNormalization(dataset) {
    // Menghitung total bobot
    const total = dataset?.criteria.reduce((acc, criterion) => acc + criterion?.weight, 0);
    
    // Menghitung hasil normalisasi masing-masing bobot
    dataset.criteria = dataset?.criteria.map(criterion => ({ ...criterion, normalizedWeight: criterion.weight / total }));
}

// function untuk menormalisasi nilai matriks
function matrixNormalization(dataset) {
    // menentukan nilai pembagi untuk masing-masing kriteria berdasarkan nilai-nilai alternatif
    dataset?.alternatives.forEach(alternative => {
        dataset.criteria = dataset?.criteria.map(criterion => {
            return {
                ...criterion,
                divisorValue: divisorValue(alternative[criterion?.code], criterion?.divisorValue ?? 0, criterion?.type),
            };
        });
    });

    // normalisasi matriks berdasarkan nilai pembaginya
    dataset?.alternatives.forEach(alternative => {
        dataset?.criteria.forEach(criterion => {
            alternative[`normalized${criterion?.code}`] = normalize(alternative[criterion?.code], criterion?.divisorValue, criterion?.type);
        });
    });
}

// function untuk menghitung pembobotan matriks yang telah dinormalisasi
function resultCalculation(dataset) {
    dataset?.alternatives.forEach(alternative => {
        alternative.result = 0;
        dataset?.criteria.forEach(criterion => {
            alternative.result += (alternative[`normalized${criterion?.code}`] * criterion?.normalizedWeight);
        });
    });
}

// function untuk menghitung nilai normalisasi matriks
function normalize(matrix, divisor, type) {
    if (type === 'cost') {
        return divisor / matrix;
    }

    return matrix / divisor;
}

// function untuk mendapatkan pembagi yang akan digunakan untuk normalisasi matriks
function divisorValue(value, initial, type) {
    if (type === 'cost' && (value < initial || initial === 0)) {
        return value;
    }

    if (type === 'benefit' && value > initial) {
        return value;
    }

    return initial;
}