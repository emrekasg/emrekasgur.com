[PostLink] = example
[PostTitle] = AWS Privatelink ile MongoDB Atlas kullanımı PoC
[Brief] = MongoDB Atlas ile AWS PrivateLink'in nasıl kullanıldığını gösteren bir PoC
[Language] = tr

### Giriş
Bu Proof of Concept (PoC), MongoDB Atlas ile AWS PrivateLink entegrasyonunu göstermektedir. Bu kurulum, AWS PrivateLink teknolojisini kullanarak, bir AWS VPC'den MongoDB Atlas'a güvenli ve özel bir bağlantı sağlar.


### Neden?
AWS PrivateLink, AWS hizmetleri ile yerel ağlar arasında özel bir bağlantı çözümü sunar. Özellikle veri güvenliği ve ağ performansının önemli olduğu durumlarda avantajlıdır. Geleneksel yöntemlere (örneğin, halka açık uç noktalar veya VPC eşleme) kıyasla, PrivateLink daha güvenli ve ölçeklenebilir bir yaklaşım sunar. Bunun nedeni, AWS ve MongoDB Atlas arasındaki trafiğin genel internet üzerinden geçmemesi, böylece potansiyel tehditlere maruz kalma ihtimalini azaltmasıdır.


### Ne zaman VPC peering yerine Privatelink kullanılmalıdır?
AWS PrivateLink, aşağıdaki durumlarda VPC eşlemeden tercih edilir:

Gelişmiş güvenlik ve gizlilik kritik önem taşıdığında.
Birden fazla VPC'nin MongoDB Atlas'a karmaşık yönlendirme veya IP çakışması sorunları olmadan bağlanması gerektiğinde.
Mimari, tutarlı düşük gecikme süreli bağlantıyı talep ettiğinde.
Organizasyon, daha az yük ile akıcı bir ağ yönetimi sürecini tercih ettiğinde.

## Deployment

#### Ön koşullar

- AWS hesabı
- MongoDB Atlas hesabı
- Terraform


Bu PoC'yi deploy etmek için aşağıdaki adımları izleyin:

1. Terraform'u aşağıdaki komutu çalıştırarak başlatın:

```sh
terraform init
```

2. Terraform'un uygulayacağı değişiklikleri aşağıdaki komutu çalıştırarak gözden geçirin:

```sh
terraform plan
```

3. Ve son komutuda çalıştırarak deploy edin
```sh
terraform apply -auto-approve
```

### GitHub Repo:

[Kaynak kodu](https://github.com/emrekasg/privatelink-atlas-poc/)
