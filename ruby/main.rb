require 'json'

# function untuk menampung kode utama, agar lebih rapi
def main
    # membaca isi file dataset.json dan menampungnya ke dalam variable file
    file = File.read 'dataset.json'

    # mengubah tipe data variable file menjadi hash dan menampungnya di variable dataset
    dataset = JSON.parse file

    # normalisasi nilai bobot
    weight_normalization dataset

    # normalisasi nilai matriks
    matrix_normalization dataset

    # menghitung pembobotan matriks yang telah dinormalisasi
    result_calculation dataset

    # mengurutkan hasil berdasarkan pembobotan matriks tertinggi
    dataset['alternatives'] = dataset['alternatives'].sort_by { |alternative| -alternative['result'] }

    # menampilkan hasil menggunakan std out CLI
    puts '================================'
    puts "No.\tNama \t\tHasil"
    puts '================================'
    dataset['alternatives'].each_with_index do |alternative, index|
        puts "%2d\t%-10s \t %.2f" % [index+1, alternative['name'], alternative['result'].round(2)]
    end
end

# function untuk menormalisasi nilai bobot
def weight_normalization(dataset)
    total = 0.0
    dataset['criteria'].each do |criterion|
        total += (criterion['weight']).to_f
    end

    dataset['criteria'].each do |criterion|
        criterion['normalized_weight'] = (criterion['weight'].to_f / total).to_f
    end
end

# function untuk menormalisasi nilai matriks
def matrix_normalization(dataset)
    # menentukan nilai pembagi untuk masing-masing kriteria berdasarkan nilai-nilai alternatif
    dataset['alternatives'].each do |alternative|
        dataset['criteria'] = dataset['criteria'].map do |criterion|
            criterion.merge('divisor_value' => divisor_value(alternative[criterion['code']], criterion['divisor_value'] || 0, criterion['type']))
        end
    end

    # normalisasi matriks berdasarkan nilai pembaginya
    dataset['alternatives'].each_with_index do |alternative, index|
        dataset['criteria'].each do |criterion|
            dataset['alternatives'][index]['normalized_' + criterion['code']] = normalize(alternative[criterion['code']], criterion['divisor_value'], criterion['type'])
        end
    end
end

# function untuk menghitung pembobotan matriks yang telah dinormalisasi
def result_calculation(dataset)
    dataset['alternatives'].each_with_index do |alternative, index|
        dataset['alternatives'][index]['result'] = 0
        dataset['criteria'].each do |criterion|
            dataset['alternatives'][index]['result'] += (alternative['normalized_' + criterion['code']] * criterion['normalized_weight']).to_f
        end
    end
end

# function untuk menghitung nilai normalisasi matriks
def normalize(matrix, divisor, type)
    return (divisor / matrix).to_f if type == 'cost'
    
    (matrix / divisor).to_f
end

# function untuk mendapatkan pembagi yang akan digunakan untuk normalisasi matriks
def divisor_value(value, initial, type)
    return value.to_f if type == 'cost' and (value < initial or initial == 0)

    return value.to_f if type == 'benefit' and value > initial

    initial.to_f
end

# memanggil function main untuk memulai menjalankan program
main
