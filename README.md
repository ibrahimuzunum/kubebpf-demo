# kubebpf-demo

Bu dal, ana `kubebpf` uygulaması değildir. Ana uygulamayı geliştirmeye başlamadan önce çalışma ortamının temel Linux gereksinimlerini kontrol eden küçük bir başlangıç aracıdır.

Şimdilik yalnızca `doctor` komutunu içerir; gerçek eBPF ve Kubernetes özellikleri sonraki geliştirme aşamalarında eklenecektir.

## Gereksinimler

- Go 1.22 veya daha yeni bir sürüm
- `doctor` kontrollerinin tamamını geçmek için Linux

## Doctor komutu

Bilgisayarın ileride yapılacak eBPF çalışmalarına uygun temel Linux özelliklerini kontrol eder:

```sh
go run ./cmd/kubebpf doctor
```

Her kontrolün sonucu `PASS` veya `FAIL` olarak gösterilir. Kontrollerden biri başarısız olursa program başarısız çıkış koduyla kapanır.

## Make hedefleri

Programı `bin/kubebpf` konumuna derlemek için:

```sh
make build
```

Doctor komutunu çalıştırmak için:

```sh
make run
```

Tüm Go paketlerinin testlerini çalıştırmak için:

```sh
make test
```
