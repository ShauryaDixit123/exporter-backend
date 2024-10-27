-- +goose Up
-- +goose StatementBegin
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('ALL', 'Albania Lek', 'Lek', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Albania Lek', symbol = 'Lek', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('AFN', 'Afghanistan Afghani', '؋', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Afghanistan Afghani', symbol = '؋', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('ARS', 'Argentina Peso', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Argentina Peso', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('AWG', 'Aruba Guilder', 'ƒ', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Aruba Guilder', symbol = 'ƒ', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('AUD', 'Australia Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Australia Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('AZN', 'Azerbaijan Manat', '₼', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Azerbaijan Manat', symbol = '₼', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BSD', 'Bahamas Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bahamas Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BBD', 'Barbados Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Barbados Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BYN', 'Belarus Ruble', 'Br', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Belarus Ruble', symbol = 'Br', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BZD', 'Belize Dollar', 'BZ$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Belize Dollar', symbol = 'BZ$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BMD', 'Bermuda Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bermuda Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BOB', 'Bolivia Bolíviano', '$b', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bolivia Bolíviano', symbol = '$b', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BAM', 'Bosnia and Herzegovina Convertible Mark', 'KM', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bosnia and Herzegovina Convertible Mark', symbol = 'KM', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BWP', 'Botswana Pula', 'P', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Botswana Pula', symbol = 'P', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BGN', 'Bulgaria Lev', 'лв', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bulgaria Lev', symbol = 'лв', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BRL', 'Brazil Real', 'R$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Brazil Real', symbol = 'R$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('BND', 'Brunei Darussalam Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Brunei Darussalam Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('KHR', 'Cambodia Riel', '៛', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cambodia Riel', symbol = '៛', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('CAD', 'Canada Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Canada Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('KYD', 'Cayman Islands Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cayman Islands Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('CLP', 'Chile Peso', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Chile Peso', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('CNY', 'China Yuan Renminbi', '¥', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'China Yuan Renminbi', symbol = '¥', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('COP', 'Colombia Peso', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Colombia Peso', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('CRC', 'Costa Rica Colon', '₡', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Costa Rica Colon', symbol = '₡', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('HRK', 'Croatia Kuna', 'kn', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Croatia Kuna', symbol = 'kn', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('CUP', 'Cuba Peso', '₱', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cuba Peso', symbol = '₱', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('CZK', 'Czech Republic Koruna', 'Kč', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Czech Republic Koruna', symbol = 'Kč', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('DKK', 'Denmark Krone', 'kr', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Denmark Krone', symbol = 'kr', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('DOP', 'Dominican Republic Peso', 'RD$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Dominican Republic Peso', symbol = 'RD$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('XCD', 'East Caribbean Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'East Caribbean Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('EGP', 'Egypt Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Egypt Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SVC', 'El Salvador Colon', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'El Salvador Colon', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('EUR', 'Euro Member Countries', '€', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Euro Member Countries', symbol = '€', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('FKP', 'Falkland Islands (Malvinas) Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Falkland Islands (Malvinas) Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('FJD', 'Fiji Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Fiji Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('GHS', 'Ghana Cedi', '¢', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ghana Cedi', symbol = '¢', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('GIP', 'Gibraltar Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Gibraltar Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('GTQ', 'Guatemala Quetzal', 'Q', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guatemala Quetzal', symbol = 'Q', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('GGP', 'Guernsey Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guernsey Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('GYD', 'Guyana Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guyana Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('HNL', 'Honduras Lempira', 'L', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Honduras Lempira', symbol = 'L', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('HKD', 'Hong Kong Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Hong Kong Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('HUF', 'Hungary Forint', 'Ft', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Hungary Forint', symbol = 'Ft', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('ISK', 'Iceland Krona', 'kr', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Iceland Krona', symbol = 'kr', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('INR', 'India Rupee', '₹', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'India Rupee', symbol = '₹', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('IDR', 'Indonesia Rupiah', 'Rp', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Indonesia Rupiah', symbol = 'Rp', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('IRR', 'Iran Rial', '﷼', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Iran Rial', symbol = '﷼', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('IMP', 'Isle of Man Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Isle of Man Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('ILS', 'Israel Shekel', '₪', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Israel Shekel', symbol = '₪', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('JMD', 'Jamaica Dollar', 'J$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Jamaica Dollar', symbol = 'J$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('JPY', 'Japan Yen', '¥', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Japan Yen', symbol = '¥', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('JEP', 'Jersey Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Jersey Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('KZT', 'Kazakhstan Tenge', 'лв', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kazakhstan Tenge', symbol = 'лв', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('KPW', 'Korea (North) Won', '₩', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Korea (North) Won', symbol = '₩', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('KRW', 'Korea (South) Won', '₩', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Korea (South) Won', symbol = '₩', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('KGS', 'Kyrgyzstan Som', 'лв', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kyrgyzstan Som', symbol = 'лв', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('LAK', 'Laos Kip', '₭', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Laos Kip', symbol = '₭', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('LBP', 'Lebanon Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Lebanon Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('LRD', 'Liberia Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Liberia Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('MKD', 'Macedonia Denar', 'ден', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Macedonia Denar', symbol = 'ден', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('MYR', 'Malaysia Ringgit', 'RM', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Malaysia Ringgit', symbol = 'RM', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('MUR', 'Mauritius Rupee', '₨', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mauritius Rupee', symbol = '₨', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('MXN', 'Mexico Peso', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mexico Peso', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('MNT', 'Mongolia Tughrik', '₮', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mongolia Tughrik', symbol = '₮', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('MNT', 'Moroccan-dirham', '&nbsp;د.إ', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Moroccan-dirham', symbol = '&nbsp;د.إ', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('MZN', 'Mozambique Metical', 'MT', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mozambique Metical', symbol = 'MT', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('NAD', 'Namibia Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Namibia Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('NPR', 'Nepal Rupee', '₨', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Nepal Rupee', symbol = '₨', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('ANG', 'Netherlands Antilles Guilder', 'ƒ', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Netherlands Antilles Guilder', symbol = 'ƒ', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('NZD', 'New Zealand Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'New Zealand Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('NIO', 'Nicaragua Cordoba', 'C$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Nicaragua Cordoba', symbol = 'C$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('NGN', 'Nigeria Naira', '₦', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Nigeria Naira', symbol = '₦', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('NOK', 'Norway Krone', 'kr', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Norway Krone', symbol = 'kr', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('OMR', 'Oman Rial', '﷼', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Oman Rial', symbol = '﷼', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('PKR', 'Pakistan Rupee', '₨', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Pakistan Rupee', symbol = '₨', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('PAB', 'Panama Balboa', 'B/.', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Panama Balboa', symbol = 'B/.', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('PYG', 'Paraguay Guarani', 'Gs', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Paraguay Guarani', symbol = 'Gs', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('PEN', 'Peru Sol', 'S/.', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Peru Sol', symbol = 'S/.', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('PHP', 'Philippines Peso', '₱', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Philippines Peso', symbol = '₱', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('PLN', 'Poland Zloty', 'zł', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Poland Zloty', symbol = 'zł', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('QAR', 'Qatar Riyal', '﷼', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Qatar Riyal', symbol = '﷼', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('RON', 'Romania Leu', 'lei', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Romania Leu', symbol = 'lei', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('RUB', 'Russia Ruble', '₽', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Russia Ruble', symbol = '₽', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SHP', 'Saint Helena Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Helena Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SAR', 'Saudi Arabia Riyal', '﷼', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saudi Arabia Riyal', symbol = '﷼', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('RSD', 'Serbia Dinar', 'Дин.', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Serbia Dinar', symbol = 'Дин.', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SCR', 'Seychelles Rupee', '₨', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Seychelles Rupee', symbol = '₨', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SGD', 'Singapore Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Singapore Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SBD', 'Solomon Islands Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Solomon Islands Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SOS', 'Somalia Shilling', 'S', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Somalia Shilling', symbol = 'S', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('KRW', 'South Korean Won', '₩', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'South Korean Won', symbol = '₩', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('ZAR', 'South Africa Rand', 'R', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'South Africa Rand', symbol = 'R', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('LKR', 'Sri Lanka Rupee', '₨', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sri Lanka Rupee', symbol = '₨', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SEK', 'Sweden Krona', 'kr', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sweden Krona', symbol = 'kr', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('CHF', 'Switzerland Franc', 'CHF', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Switzerland Franc', symbol = 'CHF', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SRD', 'Suriname Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Suriname Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('SYP', 'Syria Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Syria Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('TWD', 'Taiwan New Dollar', 'NT$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Taiwan New Dollar', symbol = 'NT$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('THB', 'Thailand Baht', '฿', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Thailand Baht', symbol = '฿', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('TTD', 'Trinidad and Tobago Dollar', 'TT$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Trinidad and Tobago Dollar', symbol = 'TT$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('TRY', 'Turkey Lira', '₺', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Turkey Lira', symbol = '₺', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('TVD', 'Tuvalu Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tuvalu Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('UAH', 'Ukraine Hryvnia', '₴', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ukraine Hryvnia', symbol = '₴', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('AED', 'UAE-Dirham', '&nbsp;د.إ', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'UAE-Dirham', symbol = '&nbsp;د.إ', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('GBP', 'United Kingdom Pound', '£', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'United Kingdom Pound', symbol = '£', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('USD', 'United States Dollar', '$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'United States Dollar', symbol = '$', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('UYU', 'Uruguay Peso', '$U', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Uruguay Peso', symbol = '$U', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('UZS', 'Uzbekistan Som', 'лв', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Uzbekistan Som', symbol = 'лв', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('VEF', 'Venezuela Bolívar', 'Bs', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Venezuela Bolívar', symbol = 'Bs', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('VND', 'Viet Nam Dong', '₫', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Viet Nam Dong', symbol = '₫', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('YER', 'Yemen Rial', '﷼', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Yemen Rial', symbol = '﷼', is_active = true;
INSERT INTO currencies (iso_code, name, symbol, is_active) VALUES ('ZWD', 'Zimbabwe Dollar', 'Z$', true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Zimbabwe Dollar', symbol = 'Z$', is_active = true;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM currencies WHERE iso_code = 'ALL';
DELETE FROM currencies WHERE iso_code = 'AFN';
DELETE FROM currencies WHERE iso_code = 'ARS';
DELETE FROM currencies WHERE iso_code = 'AWG';
DELETE FROM currencies WHERE iso_code = 'AUD';
DELETE FROM currencies WHERE iso_code = 'AZN';
DELETE FROM currencies WHERE iso_code = 'BSD';
DELETE FROM currencies WHERE iso_code = 'BBD';
DELETE FROM currencies WHERE iso_code = 'BYN';
DELETE FROM currencies WHERE iso_code = 'BZD';
DELETE FROM currencies WHERE iso_code = 'BMD';
DELETE FROM currencies WHERE iso_code = 'BOB';
DELETE FROM currencies WHERE iso_code = 'BAM';
DELETE FROM currencies WHERE iso_code = 'BWP';
DELETE FROM currencies WHERE iso_code = 'BGN';
DELETE FROM currencies WHERE iso_code = 'BRL';
DELETE FROM currencies WHERE iso_code = 'BND';
DELETE FROM currencies WHERE iso_code = 'KHR';
DELETE FROM currencies WHERE iso_code = 'CAD';
DELETE FROM currencies WHERE iso_code = 'KYD';
DELETE FROM currencies WHERE iso_code = 'CLP';
DELETE FROM currencies WHERE iso_code = 'CNY';
DELETE FROM currencies WHERE iso_code = 'COP';
DELETE FROM currencies WHERE iso_code = 'CRC';
DELETE FROM currencies WHERE iso_code = 'HRK';
DELETE FROM currencies WHERE iso_code = 'CUP';
DELETE FROM currencies WHERE iso_code = 'CZK';
DELETE FROM currencies WHERE iso_code = 'DKK';
DELETE FROM currencies WHERE iso_code = 'DOP';
DELETE FROM currencies WHERE iso_code = 'XCD';
DELETE FROM currencies WHERE iso_code = 'EGP';
DELETE FROM currencies WHERE iso_code = 'SVC';
DELETE FROM currencies WHERE iso_code = 'EUR';
DELETE FROM currencies WHERE iso_code = 'FKP';
DELETE FROM currencies WHERE iso_code = 'FJD';
DELETE FROM currencies WHERE iso_code = 'GHS';
DELETE FROM currencies WHERE iso_code = 'GIP';
DELETE FROM currencies WHERE iso_code = 'GTQ';
DELETE FROM currencies WHERE iso_code = 'GGP';
DELETE FROM currencies WHERE iso_code = 'GYD';
DELETE FROM currencies WHERE iso_code = 'HNL';
DELETE FROM currencies WHERE iso_code = 'HKD';
DELETE FROM currencies WHERE iso_code = 'HUF';
DELETE FROM currencies WHERE iso_code = 'ISK';
DELETE FROM currencies WHERE iso_code = 'INR';
DELETE FROM currencies WHERE iso_code = 'IDR';
DELETE FROM currencies WHERE iso_code = 'IRR';
DELETE FROM currencies WHERE iso_code = 'IMP';
DELETE FROM currencies WHERE iso_code = 'ILS';
DELETE FROM currencies WHERE iso_code = 'JMD';
DELETE FROM currencies WHERE iso_code = 'JPY';
DELETE FROM currencies WHERE iso_code = 'JEP';
DELETE FROM currencies WHERE iso_code = 'KZT';
DELETE FROM currencies WHERE iso_code = 'KPW';
DELETE FROM currencies WHERE iso_code = 'KRW';
DELETE FROM currencies WHERE iso_code = 'KGS';
DELETE FROM currencies WHERE iso_code = 'LAK';
DELETE FROM currencies WHERE iso_code = 'LBP';
DELETE FROM currencies WHERE iso_code = 'LRD';
DELETE FROM currencies WHERE iso_code = 'MKD';
DELETE FROM currencies WHERE iso_code = 'MYR';
DELETE FROM currencies WHERE iso_code = 'MUR';
DELETE FROM currencies WHERE iso_code = 'MXN';
DELETE FROM currencies WHERE iso_code = 'MNT';
DELETE FROM currencies WHERE iso_code = 'MNT';
DELETE FROM currencies WHERE iso_code = 'MZN';
DELETE FROM currencies WHERE iso_code = 'NAD';
DELETE FROM currencies WHERE iso_code = 'NPR';
DELETE FROM currencies WHERE iso_code = 'ANG';
DELETE FROM currencies WHERE iso_code = 'NZD';
DELETE FROM currencies WHERE iso_code = 'NIO';
DELETE FROM currencies WHERE iso_code = 'NGN';
DELETE FROM currencies WHERE iso_code = 'NOK';
DELETE FROM currencies WHERE iso_code = 'OMR';
DELETE FROM currencies WHERE iso_code = 'PKR';
DELETE FROM currencies WHERE iso_code = 'PAB';
DELETE FROM currencies WHERE iso_code = 'PYG';
DELETE FROM currencies WHERE iso_code = 'PEN';
DELETE FROM currencies WHERE iso_code = 'PHP';
DELETE FROM currencies WHERE iso_code = 'PLN';
DELETE FROM currencies WHERE iso_code = 'QAR';
DELETE FROM currencies WHERE iso_code = 'RON';
DELETE FROM currencies WHERE iso_code = 'RUB';
DELETE FROM currencies WHERE iso_code = 'SHP';
DELETE FROM currencies WHERE iso_code = 'SAR';
DELETE FROM currencies WHERE iso_code = 'RSD';
DELETE FROM currencies WHERE iso_code = 'SCR';
DELETE FROM currencies WHERE iso_code = 'SGD';
DELETE FROM currencies WHERE iso_code = 'SBD';
DELETE FROM currencies WHERE iso_code = 'SOS';
DELETE FROM currencies WHERE iso_code = 'KRW';
DELETE FROM currencies WHERE iso_code = 'ZAR';
DELETE FROM currencies WHERE iso_code = 'LKR';
DELETE FROM currencies WHERE iso_code = 'SEK';
DELETE FROM currencies WHERE iso_code = 'CHF';
DELETE FROM currencies WHERE iso_code = 'SRD';
DELETE FROM currencies WHERE iso_code = 'SYP';
DELETE FROM currencies WHERE iso_code = 'TWD';
DELETE FROM currencies WHERE iso_code = 'THB';
DELETE FROM currencies WHERE iso_code = 'TTD';
DELETE FROM currencies WHERE iso_code = 'TRY';
DELETE FROM currencies WHERE iso_code = 'TVD';
DELETE FROM currencies WHERE iso_code = 'UAH';
DELETE FROM currencies WHERE iso_code = 'AED';
DELETE FROM currencies WHERE iso_code = 'GBP';
DELETE FROM currencies WHERE iso_code = 'USD';
DELETE FROM currencies WHERE iso_code = 'UYU';
DELETE FROM currencies WHERE iso_code = 'UZS';
DELETE FROM currencies WHERE iso_code = 'VEF';
DELETE FROM currencies WHERE iso_code = 'VND';
DELETE FROM currencies WHERE iso_code = 'YER';
DELETE FROM currencies WHERE iso_code = 'ZWD';

-- +goose StatementEnd
