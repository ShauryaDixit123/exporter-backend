-- +goose Up
-- +goose StatementBegin
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AC', 'Ascension Island', 'ACS', 247, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ascension Island', iso_3 = 'ACS', calling_code = 247, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AD', 'Andorra', 'AND', 376, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Andorra', iso_3 = 'AND', calling_code = 376, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AE', 'United Arab Emirates', 'ARE', 971, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'United Arab Emirates', iso_3 = 'ARE', calling_code = 971, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AF', 'Afghanistan', 'AFG', 93, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Afghanistan', iso_3 = 'AFG', calling_code = 93, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AG', 'Antigua and Barbuda', 'ATG', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Antigua and Barbuda', iso_3 = 'ATG', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AI', 'Anguilla', 'AIA', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Anguilla', iso_3 = 'AIA', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AL', 'Albania', 'ALB', 355, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Albania', iso_3 = 'ALB', calling_code = 355, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AM', 'Armenia', 'ARM', 374, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Armenia', iso_3 = 'ARM', calling_code = 374, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AO', 'Angola', 'AGO', 244, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Angola', iso_3 = 'AGO', calling_code = 244, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AR', 'Argentina', 'ARG', 54, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Argentina', iso_3 = 'ARG', calling_code = 54, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AS', 'American Samoa', 'ASM', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'American Samoa', iso_3 = 'ASM', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AT', 'Austria', 'AUT', 43, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Austria', iso_3 = 'AUT', calling_code = 43, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AU', 'Australia', 'AUS', 61, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Australia', iso_3 = 'AUS', calling_code = 61, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AW', 'Aruba', 'ABW', 297, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Aruba', iso_3 = 'ABW', calling_code = 297, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AX', 'Åland Islands', 'ALA', 358, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Åland Islands', iso_3 = 'ALA', calling_code = 358, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('AZ', 'Azerbaijan', 'AZE', 994, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Azerbaijan', iso_3 = 'AZE', calling_code = 994, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BA', 'Bosnia and Herzegovina', 'BIH', 387, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bosnia and Herzegovina', iso_3 = 'BIH', calling_code = 387, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BB', 'Barbados', 'BRB', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Barbados', iso_3 = 'BRB', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BD', 'Bangladesh', 'BGD', 880, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bangladesh', iso_3 = 'BGD', calling_code = 880, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BE', 'Belgium', 'BEL', 32, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Belgium', iso_3 = 'BEL', calling_code = 32, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BF', 'Burkina Faso', 'BFA', 226, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Burkina Faso', iso_3 = 'BFA', calling_code = 226, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BG', 'Bulgaria', 'BGR', 359, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bulgaria', iso_3 = 'BGR', calling_code = 359, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BH', 'Bahrain', 'BHR', 973, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bahrain', iso_3 = 'BHR', calling_code = 973, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BI', 'Burundi', 'BDI', 257, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Burundi', iso_3 = 'BDI', calling_code = 257, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BJ', 'Benin', 'BEN', 229, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Benin', iso_3 = 'BEN', calling_code = 229, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BL', 'Saint Barthélemy', 'BLM', 590, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Barthélemy', iso_3 = 'BLM', calling_code = 590, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BM', 'Bermuda', 'BMU', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bermuda', iso_3 = 'BMU', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BN', 'Brunei', 'BRN', 673, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Brunei', iso_3 = 'BRN', calling_code = 673, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BO', 'Bolivia', 'BOL', 591, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bolivia', iso_3 = 'BOL', calling_code = 591, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BQ', 'Bonaire, Sint Eustatius and Saba', 'BES', 599, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bonaire, Sint Eustatius and Saba', iso_3 = 'BES', calling_code = 599, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BR', 'Brazil', 'BRA', 55, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Brazil', iso_3 = 'BRA', calling_code = 55, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BS', 'Bahamas', 'BHS', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bahamas', iso_3 = 'BHS', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BT', 'Bhutan', 'BTN', 975, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Bhutan', iso_3 = 'BTN', calling_code = 975, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BW', 'Botswana', 'BWA', 267, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Botswana', iso_3 = 'BWA', calling_code = 267, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BY', 'Belarus', 'BLR', 375, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Belarus', iso_3 = 'BLR', calling_code = 375, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('BZ', 'Belize', 'BLZ', 501, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Belize', iso_3 = 'BLZ', calling_code = 501, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CA', 'Canada', 'CAN', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Canada', iso_3 = 'CAN', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CC', 'Cocos Islands', 'CCK', 61, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cocos Islands', iso_3 = 'CCK', calling_code = 61, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CD', 'The Democratic Republic Of Congo', 'COD', 243, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'The Democratic Republic Of Congo', iso_3 = 'COD', calling_code = 243, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CF', 'Central African Republic', 'CAF', 236, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Central African Republic', iso_3 = 'CAF', calling_code = 236, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CG', 'Congo', 'COG', 242, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Congo', iso_3 = 'COG', calling_code = 242, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CH', 'Switzerland', 'CHE', 41, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Switzerland', iso_3 = 'CHE', calling_code = 41, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CI', 'Côte d''Ivoire', 'CIV', 225, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Côte d''Ivoire', iso_3 = 'CIV', calling_code = 225, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CK', 'Cook Islands', 'COK', 682, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cook Islands', iso_3 = 'COK', calling_code = 682, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CL', 'Chile', 'CHL', 56, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Chile', iso_3 = 'CHL', calling_code = 56, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CM', 'Cameroon', 'CMR', 237, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cameroon', iso_3 = 'CMR', calling_code = 237, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CN', 'China', 'CHN', 86, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'China', iso_3 = 'CHN', calling_code = 86, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CO', 'Colombia', 'COL', 57, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Colombia', iso_3 = 'COL', calling_code = 57, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CR', 'Costa Rica', 'CRI', 506, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Costa Rica', iso_3 = 'CRI', calling_code = 506, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CU', 'Cuba', 'CUB', 53, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cuba', iso_3 = 'CUB', calling_code = 53, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CV', 'Cape Verde', 'CPV', 238, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cape Verde', iso_3 = 'CPV', calling_code = 238, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CW', 'Curaçao', 'CUW', 599, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Curaçao', iso_3 = 'CUW', calling_code = 599, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CX', 'Christmas Island', 'CXR', 61, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Christmas Island', iso_3 = 'CXR', calling_code = 61, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CY', 'Cyprus', 'CYP', 357, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cyprus', iso_3 = 'CYP', calling_code = 357, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('CZ', 'Czech Republic', 'CZE', 420, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Czech Republic', iso_3 = 'CZE', calling_code = 420, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('DE', 'Germany', 'DEU', 49, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Germany', iso_3 = 'DEU', calling_code = 49, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('DJ', 'Djibouti', 'DJI', 253, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Djibouti', iso_3 = 'DJI', calling_code = 253, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('DK', 'Denmark', 'DNK', 45, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Denmark', iso_3 = 'DNK', calling_code = 45, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('DM', 'Dominica', 'DMA', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Dominica', iso_3 = 'DMA', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('DO', 'Dominican Republic', 'DOM', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Dominican Republic', iso_3 = 'DOM', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('DZ', 'Algeria', 'DZA', 213, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Algeria', iso_3 = 'DZA', calling_code = 213, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('EC', 'Ecuador', 'ECU', 593, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ecuador', iso_3 = 'ECU', calling_code = 593, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('EE', 'Estonia', 'EST', 372, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Estonia', iso_3 = 'EST', calling_code = 372, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('EG', 'Egypt', 'EGY', 20, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Egypt', iso_3 = 'EGY', calling_code = 20, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('EH', 'Western Sahara', 'ESH', 212, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Western Sahara', iso_3 = 'ESH', calling_code = 212, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ER', 'Eritrea', 'ERI', 291, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Eritrea', iso_3 = 'ERI', calling_code = 291, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ES', 'Spain', 'ESP', 34, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Spain', iso_3 = 'ESP', calling_code = 34, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ET', 'Ethiopia', 'ETH', 251, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ethiopia', iso_3 = 'ETH', calling_code = 251, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('FI', 'Finland', 'FIN', 358, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Finland', iso_3 = 'FIN', calling_code = 358, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('FJ', 'Fiji', 'FJI', 679, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Fiji', iso_3 = 'FJI', calling_code = 679, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('FK', 'Falkland Islands', 'FLK', 500, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Falkland Islands', iso_3 = 'FLK', calling_code = 500, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('FM', 'Micronesia', 'FSM', 691, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Micronesia', iso_3 = 'FSM', calling_code = 691, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('FO', 'Faroe Islands', 'FRO', 298, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Faroe Islands', iso_3 = 'FRO', calling_code = 298, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('FR', 'France', 'FRA', 33, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'France', iso_3 = 'FRA', calling_code = 33, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GA', 'Gabon', 'GAB', 241, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Gabon', iso_3 = 'GAB', calling_code = 241, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GB', 'United Kingdom', 'GBR', 44, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'United Kingdom', iso_3 = 'GBR', calling_code = 44, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GD', 'Grenada', 'GRD', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Grenada', iso_3 = 'GRD', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GE', 'Georgia', 'GEO', 995, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Georgia', iso_3 = 'GEO', calling_code = 995, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GF', 'French Guiana', 'GUF', 594, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'French Guiana', iso_3 = 'GUF', calling_code = 594, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GG', 'Guernsey', 'GGY', 44, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guernsey', iso_3 = 'GGY', calling_code = 44, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GH', 'Ghana', 'GHA', 233, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ghana', iso_3 = 'GHA', calling_code = 233, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GI', 'Gibraltar', 'GIB', 350, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Gibraltar', iso_3 = 'GIB', calling_code = 350, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GL', 'Greenland', 'GRL', 299, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Greenland', iso_3 = 'GRL', calling_code = 299, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GM', 'Gambia', 'GMB', 220, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Gambia', iso_3 = 'GMB', calling_code = 220, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GN', 'Guinea', 'GIN', 224, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guinea', iso_3 = 'GIN', calling_code = 224, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GP', 'Guadeloupe', 'GLP', 590, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guadeloupe', iso_3 = 'GLP', calling_code = 590, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GQ', 'Equatorial Guinea', 'GNQ', 240, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Equatorial Guinea', iso_3 = 'GNQ', calling_code = 240, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GR', 'Greece', 'GRC', 30, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Greece', iso_3 = 'GRC', calling_code = 30, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GT', 'Guatemala', 'GTM', 502, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guatemala', iso_3 = 'GTM', calling_code = 502, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GU', 'Guam', 'GUM', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guam', iso_3 = 'GUM', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GW', 'Guinea-Bissau', 'GNB', 245, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guinea-Bissau', iso_3 = 'GNB', calling_code = 245, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('GY', 'Guyana', 'GUY', 592, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Guyana', iso_3 = 'GUY', calling_code = 592, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('HK', 'Hong Kong', 'HKG', 852, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Hong Kong', iso_3 = 'HKG', calling_code = 852, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('HN', 'Honduras', 'HND', 504, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Honduras', iso_3 = 'HND', calling_code = 504, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('HR', 'Croatia', 'HRV', 385, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Croatia', iso_3 = 'HRV', calling_code = 385, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('HT', 'Haiti', 'HTI', 509, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Haiti', iso_3 = 'HTI', calling_code = 509, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('HU', 'Hungary', 'HUN', 36, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Hungary', iso_3 = 'HUN', calling_code = 36, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ID', 'Indonesia', 'IDN', 62, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Indonesia', iso_3 = 'IDN', calling_code = 62, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IE', 'Ireland', 'IRL', 353, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ireland', iso_3 = 'IRL', calling_code = 353, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IL', 'Israel', 'ISR', 972, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Israel', iso_3 = 'ISR', calling_code = 972, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IM', 'Isle Of Man', 'IMN', 44, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Isle Of Man', iso_3 = 'IMN', calling_code = 44, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IN', 'India', 'IND', 91, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'India', iso_3 = 'IND', calling_code = 91, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IO', 'British Indian Ocean Territory', 'IOT', 246, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'British Indian Ocean Territory', iso_3 = 'IOT', calling_code = 246, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IQ', 'Iraq', 'IRQ', 964, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Iraq', iso_3 = 'IRQ', calling_code = 964, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IR', 'Iran', 'IRN', 98, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Iran', iso_3 = 'IRN', calling_code = 98, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IS', 'Iceland', 'ISL', 354, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Iceland', iso_3 = 'ISL', calling_code = 354, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('IT', 'Italy', 'ITA', 39, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Italy', iso_3 = 'ITA', calling_code = 39, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('JE', 'Jersey', 'JEY', 44, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Jersey', iso_3 = 'JEY', calling_code = 44, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('JM', 'Jamaica', 'JAM', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Jamaica', iso_3 = 'JAM', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('JO', 'Jordan', 'JOR', 962, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Jordan', iso_3 = 'JOR', calling_code = 962, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('JP', 'Japan', 'JPN', 81, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Japan', iso_3 = 'JPN', calling_code = 81, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KE', 'Kenya', 'KEN', 254, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kenya', iso_3 = 'KEN', calling_code = 254, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KG', 'Kyrgyzstan', 'KGZ', 996, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kyrgyzstan', iso_3 = 'KGZ', calling_code = 996, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KH', 'Cambodia', 'KHM', 855, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cambodia', iso_3 = 'KHM', calling_code = 855, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KI', 'Kiribati', 'KIR', 686, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kiribati', iso_3 = 'KIR', calling_code = 686, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KM', 'Comoros', 'COM', 269, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Comoros', iso_3 = 'COM', calling_code = 269, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KN', 'Saint Kitts And Nevis', 'KNA', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Kitts And Nevis', iso_3 = 'KNA', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KP', 'North Korea', 'PRK', 850, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'North Korea', iso_3 = 'PRK', calling_code = 850, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KR', 'South Korea', 'KOR', 82, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'South Korea', iso_3 = 'KOR', calling_code = 82, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KW', 'Kuwait', 'KWT', 965, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kuwait', iso_3 = 'KWT', calling_code = 965, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KY', 'Cayman Islands', 'CYM', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Cayman Islands', iso_3 = 'CYM', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('KZ', 'Kazakhstan', 'KAZ', 7, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kazakhstan', iso_3 = 'KAZ', calling_code = 7, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LA', 'Laos', 'LAO', 856, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Laos', iso_3 = 'LAO', calling_code = 856, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LB', 'Lebanon', 'LBN', 961, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Lebanon', iso_3 = 'LBN', calling_code = 961, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LC', 'Saint Lucia', 'LCA', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Lucia', iso_3 = 'LCA', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LI', 'Liechtenstein', 'LIE', 423, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Liechtenstein', iso_3 = 'LIE', calling_code = 423, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LK', 'Sri Lanka', 'LKA', 94, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sri Lanka', iso_3 = 'LKA', calling_code = 94, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LR', 'Liberia', 'LBR', 231, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Liberia', iso_3 = 'LBR', calling_code = 231, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LS', 'Lesotho', 'LSO', 266, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Lesotho', iso_3 = 'LSO', calling_code = 266, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LT', 'Lithuania', 'LTU', 370, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Lithuania', iso_3 = 'LTU', calling_code = 370, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LU', 'Luxembourg', 'LUX', 352, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Luxembourg', iso_3 = 'LUX', calling_code = 352, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LV', 'Latvia', 'LVA', 371, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Latvia', iso_3 = 'LVA', calling_code = 371, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('LY', 'Libya', 'LBY', 218, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Libya', iso_3 = 'LBY', calling_code = 218, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MA', 'Morocco', 'MAR', 212, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Morocco', iso_3 = 'MAR', calling_code = 212, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MC', 'Monaco', 'MCO', 377, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Monaco', iso_3 = 'MCO', calling_code = 377, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MD', 'Moldova', 'MDA', 373, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Moldova', iso_3 = 'MDA', calling_code = 373, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ME', 'Montenegro', 'MNE', 382, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Montenegro', iso_3 = 'MNE', calling_code = 382, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MF', 'Saint Martin', 'MAF', 590, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Martin', iso_3 = 'MAF', calling_code = 590, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MG', 'Madagascar', 'MDG', 261, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Madagascar', iso_3 = 'MDG', calling_code = 261, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MH', 'Marshall Islands', 'MHL', 692, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Marshall Islands', iso_3 = 'MHL', calling_code = 692, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MK', 'Macedonia', 'MKD', 389, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Macedonia', iso_3 = 'MKD', calling_code = 389, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ML', 'Mali', 'MLI', 223, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mali', iso_3 = 'MLI', calling_code = 223, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MM', 'Myanmar', 'MMR', 95, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Myanmar', iso_3 = 'MMR', calling_code = 95, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MN', 'Mongolia', 'MNG', 976, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mongolia', iso_3 = 'MNG', calling_code = 976, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MO', 'Macao', 'MAC', 853, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Macao', iso_3 = 'MAC', calling_code = 853, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MP', 'Northern Mariana Islands', 'MNP', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Northern Mariana Islands', iso_3 = 'MNP', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MQ', 'Martinique', 'MTQ', 596, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Martinique', iso_3 = 'MTQ', calling_code = 596, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MR', 'Mauritania', 'MRT', 222, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mauritania', iso_3 = 'MRT', calling_code = 222, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MS', 'Montserrat', 'MSR', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Montserrat', iso_3 = 'MSR', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MT', 'Malta', 'MLT', 356, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Malta', iso_3 = 'MLT', calling_code = 356, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MU', 'Mauritius', 'MUS', 230, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mauritius', iso_3 = 'MUS', calling_code = 230, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MV', 'Maldives', 'MDV', 960, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Maldives', iso_3 = 'MDV', calling_code = 960, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MW', 'Malawi', 'MWI', 265, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Malawi', iso_3 = 'MWI', calling_code = 265, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MX', 'Mexico', 'MEX', 52, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mexico', iso_3 = 'MEX', calling_code = 52, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MY', 'Malaysia', 'MYS', 60, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Malaysia', iso_3 = 'MYS', calling_code = 60, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('MZ', 'Mozambique', 'MOZ', 258, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mozambique', iso_3 = 'MOZ', calling_code = 258, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NA', 'Namibia', 'NAM', 264, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Namibia', iso_3 = 'NAM', calling_code = 264, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NC', 'New Caledonia', 'NCL', 687, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'New Caledonia', iso_3 = 'NCL', calling_code = 687, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NE', 'Niger', 'NER', 227, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Niger', iso_3 = 'NER', calling_code = 227, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NF', 'Norfolk Island', 'NFK', 672, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Norfolk Island', iso_3 = 'NFK', calling_code = 672, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NG', 'Nigeria', 'NGA', 234, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Nigeria', iso_3 = 'NGA', calling_code = 234, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NI', 'Nicaragua', 'NIC', 505, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Nicaragua', iso_3 = 'NIC', calling_code = 505, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NL', 'Netherlands', 'NLD', 31, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Netherlands', iso_3 = 'NLD', calling_code = 31, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NO', 'Norway', 'NOR', 47, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Norway', iso_3 = 'NOR', calling_code = 47, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NP', 'Nepal', 'NPL', 977, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Nepal', iso_3 = 'NPL', calling_code = 977, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NR', 'Nauru', 'NRU', 674, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Nauru', iso_3 = 'NRU', calling_code = 674, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NU', 'Niue', 'NIU', 683, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Niue', iso_3 = 'NIU', calling_code = 683, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('NZ', 'New Zealand', 'NZL', 64, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'New Zealand', iso_3 = 'NZL', calling_code = 64, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('OM', 'Oman', 'OMN', 968, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Oman', iso_3 = 'OMN', calling_code = 968, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PA', 'Panama', 'PAN', 507, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Panama', iso_3 = 'PAN', calling_code = 507, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PE', 'Peru', 'PER', 51, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Peru', iso_3 = 'PER', calling_code = 51, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PF', 'French Polynesia', 'PYF', 689, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'French Polynesia', iso_3 = 'PYF', calling_code = 689, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PG', 'Papua New Guinea', 'PNG', 675, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Papua New Guinea', iso_3 = 'PNG', calling_code = 675, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PH', 'Philippines', 'PHL', 63, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Philippines', iso_3 = 'PHL', calling_code = 63, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PK', 'Pakistan', 'PAK', 92, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Pakistan', iso_3 = 'PAK', calling_code = 92, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PL', 'Poland', 'POL', 48, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Poland', iso_3 = 'POL', calling_code = 48, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PM', 'Saint Pierre And Miquelon', 'SPM', 508, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Pierre And Miquelon', iso_3 = 'SPM', calling_code = 508, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PR', 'Puerto Rico', 'PRI', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Puerto Rico', iso_3 = 'PRI', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PS', 'Palestine', 'PSE', 970, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Palestine', iso_3 = 'PSE', calling_code = 970, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PT', 'Portugal', 'PRT', 351, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Portugal', iso_3 = 'PRT', calling_code = 351, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PW', 'Palau', 'PLW', 680, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Palau', iso_3 = 'PLW', calling_code = 680, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('PY', 'Paraguay', 'PRY', 595, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Paraguay', iso_3 = 'PRY', calling_code = 595, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('QA', 'Qatar', 'QAT', 974, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Qatar', iso_3 = 'QAT', calling_code = 974, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('RE', 'Reunion', 'REU', 262, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Reunion', iso_3 = 'REU', calling_code = 262, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('RO', 'Romania', 'ROU', 40, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Romania', iso_3 = 'ROU', calling_code = 40, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('RS', 'Serbia', 'SRB', 381, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Serbia', iso_3 = 'SRB', calling_code = 381, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('RU', 'Russia', 'RUS', 7, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Russia', iso_3 = 'RUS', calling_code = 7, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('RW', 'Rwanda', 'RWA', 250, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Rwanda', iso_3 = 'RWA', calling_code = 250, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SA', 'Saudi Arabia', 'SAU', 966, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saudi Arabia', iso_3 = 'SAU', calling_code = 966, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SB', 'Solomon Islands', 'SLB', 677, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Solomon Islands', iso_3 = 'SLB', calling_code = 677, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SC', 'Seychelles', 'SYC', 248, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Seychelles', iso_3 = 'SYC', calling_code = 248, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SD', 'Sudan', 'SDN', 249, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sudan', iso_3 = 'SDN', calling_code = 249, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SE', 'Sweden', 'SWE', 46, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sweden', iso_3 = 'SWE', calling_code = 46, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SG', 'Singapore', 'SGP', 65, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Singapore', iso_3 = 'SGP', calling_code = 65, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SH', 'Saint Helena', 'SHN', 290, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Helena', iso_3 = 'SHN', calling_code = 290, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SI', 'Slovenia', 'SVN', 386, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Slovenia', iso_3 = 'SVN', calling_code = 386, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SJ', 'Svalbard And Jan Mayen', 'SJM', 47, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Svalbard And Jan Mayen', iso_3 = 'SJM', calling_code = 47, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SK', 'Slovakia', 'SVK', 421, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Slovakia', iso_3 = 'SVK', calling_code = 421, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SL', 'Sierra Leone', 'SLE', 232, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sierra Leone', iso_3 = 'SLE', calling_code = 232, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SM', 'San Marino', 'SMR', 378, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'San Marino', iso_3 = 'SMR', calling_code = 378, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SN', 'Senegal', 'SEN', 221, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Senegal', iso_3 = 'SEN', calling_code = 221, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SO', 'Somalia', 'SOM', 252, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Somalia', iso_3 = 'SOM', calling_code = 252, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SR', 'Suriname', 'SUR', 597, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Suriname', iso_3 = 'SUR', calling_code = 597, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SS', 'South Sudan', 'SSD', 211, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'South Sudan', iso_3 = 'SSD', calling_code = 211, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ST', 'Sao Tome And Principe', 'STP', 239, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sao Tome And Principe', iso_3 = 'STP', calling_code = 239, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SV', 'El Salvador', 'SLV', 503, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'El Salvador', iso_3 = 'SLV', calling_code = 503, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SX', 'Sint Maarten (Dutch part)', 'SXM', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Sint Maarten (Dutch part)', iso_3 = 'SXM', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SY', 'Syria', 'SYR', 963, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Syria', iso_3 = 'SYR', calling_code = 963, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('SZ', 'Swaziland', 'SWZ', 268, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Swaziland', iso_3 = 'SWZ', calling_code = 268, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TA', 'Tristan da Cunha', 'NULL', 290, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tristan da Cunha', iso_3 = 'NULL', calling_code = 290, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TC', 'Turks And Caicos Islands', 'TCA', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Turks And Caicos Islands', iso_3 = 'TCA', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TD', 'Chad', 'TCD', 235, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Chad', iso_3 = 'TCD', calling_code = 235, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TG', 'Togo', 'TGO', 228, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Togo', iso_3 = 'TGO', calling_code = 228, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TH', 'Thailand', 'THA', 66, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Thailand', iso_3 = 'THA', calling_code = 66, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TJ', 'Tajikistan', 'TJK', 992, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tajikistan', iso_3 = 'TJK', calling_code = 992, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TK', 'Tokelau', 'TKL', 690, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tokelau', iso_3 = 'TKL', calling_code = 690, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TL', 'Timor-Leste', 'TLS', 670, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Timor-Leste', iso_3 = 'TLS', calling_code = 670, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TM', 'Turkmenistan', 'TKM', 993, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Turkmenistan', iso_3 = 'TKM', calling_code = 993, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TN', 'Tunisia', 'TUN', 216, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tunisia', iso_3 = 'TUN', calling_code = 216, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TO', 'Tonga', 'TON', 676, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tonga', iso_3 = 'TON', calling_code = 676, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TR', 'Turkey', 'TUR', 90, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Turkey', iso_3 = 'TUR', calling_code = 90, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TT', 'Trinidad and Tobago', 'TTO', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Trinidad and Tobago', iso_3 = 'TTO', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TV', 'Tuvalu', 'TUV', 688, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tuvalu', iso_3 = 'TUV', calling_code = 688, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TW', 'Taiwan', 'TWN', 886, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Taiwan', iso_3 = 'TWN', calling_code = 886, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('TZ', 'Tanzania', 'TZA', 255, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Tanzania', iso_3 = 'TZA', calling_code = 255, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('UA', 'Ukraine', 'UKR', 380, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Ukraine', iso_3 = 'UKR', calling_code = 380, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('UG', 'Uganda', 'UGA', 256, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Uganda', iso_3 = 'UGA', calling_code = 256, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('US', 'United States', 'USA', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'United States', iso_3 = 'USA', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('UY', 'Uruguay', 'URY', 598, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Uruguay', iso_3 = 'URY', calling_code = 598, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('UZ', 'Uzbekistan', 'UZB', 998, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Uzbekistan', iso_3 = 'UZB', calling_code = 998, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('VA', 'Vatican', 'VAT', 39, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Vatican', iso_3 = 'VAT', calling_code = 39, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('VC', 'Saint Vincent And The Grenadines', 'VCT', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Saint Vincent And The Grenadines', iso_3 = 'VCT', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('VE', 'Venezuela', 'VEN', 58, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Venezuela', iso_3 = 'VEN', calling_code = 58, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('VG', 'British Virgin Islands', 'VGB', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'British Virgin Islands', iso_3 = 'VGB', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('VI', 'U.S. Virgin Islands', 'VIR', 1, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'U.S. Virgin Islands', iso_3 = 'VIR', calling_code = 1, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('VN', 'Vietnam', 'VNM', 84, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Vietnam', iso_3 = 'VNM', calling_code = 84, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('VU', 'Vanuatu', 'VUT', 678, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Vanuatu', iso_3 = 'VUT', calling_code = 678, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('WF', 'Wallis And Futuna', 'WLF', 681, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Wallis And Futuna', iso_3 = 'WLF', calling_code = 681, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('WS', 'Samoa', 'WSM', 685, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Samoa', iso_3 = 'WSM', calling_code = 685, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('XK', 'Kosovo', 'KOS', 383, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Kosovo', iso_3 = 'KOS', calling_code = 383, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('YE', 'Yemen', 'YEM', 967, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Yemen', iso_3 = 'YEM', calling_code = 967, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('YT', 'Mayotte', 'MYT', 262, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Mayotte', iso_3 = 'MYT', calling_code = 262, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ZA', 'South Africa', 'ZAF', 27, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'South Africa', iso_3 = 'ZAF', calling_code = 27, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ZM', 'Zambia', 'ZMB', 260, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Zambia', iso_3 = 'ZMB', calling_code = 260, is_active = true;
INSERT INTO countries (iso_code, name, iso_3, calling_code, is_active) VALUES ('ZW', 'Zimbabwe', 'ZWE', 263, true)
ON CONFLICT (iso_code)
DO 
   UPDATE SET name = 'Zimbabwe', iso_3 = 'ZWE', calling_code = 263, is_active = true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM countries WHERE iso_code='AC';
DELETE FROM countries WHERE iso_code='AD';
DELETE FROM countries WHERE iso_code='AE';
DELETE FROM countries WHERE iso_code='AF';
DELETE FROM countries WHERE iso_code='AG';
DELETE FROM countries WHERE iso_code='AI';
DELETE FROM countries WHERE iso_code='AL';
DELETE FROM countries WHERE iso_code='AM';
DELETE FROM countries WHERE iso_code='AO';
DELETE FROM countries WHERE iso_code='AR';
DELETE FROM countries WHERE iso_code='AS';
DELETE FROM countries WHERE iso_code='AT';
DELETE FROM countries WHERE iso_code='AU';
DELETE FROM countries WHERE iso_code='AW';
DELETE FROM countries WHERE iso_code='AX';
DELETE FROM countries WHERE iso_code='AZ';
DELETE FROM countries WHERE iso_code='BA';
DELETE FROM countries WHERE iso_code='BB';
DELETE FROM countries WHERE iso_code='BD';
DELETE FROM countries WHERE iso_code='BE';
DELETE FROM countries WHERE iso_code='BF';
DELETE FROM countries WHERE iso_code='BG';
DELETE FROM countries WHERE iso_code='BH';
DELETE FROM countries WHERE iso_code='BI';
DELETE FROM countries WHERE iso_code='BJ';
DELETE FROM countries WHERE iso_code='BL';
DELETE FROM countries WHERE iso_code='BM';
DELETE FROM countries WHERE iso_code='BN';
DELETE FROM countries WHERE iso_code='BO';
DELETE FROM countries WHERE iso_code='BQ';
DELETE FROM countries WHERE iso_code='BR';
DELETE FROM countries WHERE iso_code='BS';
DELETE FROM countries WHERE iso_code='BT';
DELETE FROM countries WHERE iso_code='BW';
DELETE FROM countries WHERE iso_code='BY';
DELETE FROM countries WHERE iso_code='BZ';
DELETE FROM countries WHERE iso_code='CA';
DELETE FROM countries WHERE iso_code='CC';
DELETE FROM countries WHERE iso_code='CD';
DELETE FROM countries WHERE iso_code='CF';
DELETE FROM countries WHERE iso_code='CG';
DELETE FROM countries WHERE iso_code='CH';
DELETE FROM countries WHERE iso_code='CI';
DELETE FROM countries WHERE iso_code='CK';
DELETE FROM countries WHERE iso_code='CL';
DELETE FROM countries WHERE iso_code='CM';
DELETE FROM countries WHERE iso_code='CN';
DELETE FROM countries WHERE iso_code='CO';
DELETE FROM countries WHERE iso_code='CR';
DELETE FROM countries WHERE iso_code='CU';
DELETE FROM countries WHERE iso_code='CV';
DELETE FROM countries WHERE iso_code='CW';
DELETE FROM countries WHERE iso_code='CX';
DELETE FROM countries WHERE iso_code='CY';
DELETE FROM countries WHERE iso_code='CZ';
DELETE FROM countries WHERE iso_code='DE';
DELETE FROM countries WHERE iso_code='DJ';
DELETE FROM countries WHERE iso_code='DK';
DELETE FROM countries WHERE iso_code='DM';
DELETE FROM countries WHERE iso_code='DO';
DELETE FROM countries WHERE iso_code='DZ';
DELETE FROM countries WHERE iso_code='EC';
DELETE FROM countries WHERE iso_code='EE';
DELETE FROM countries WHERE iso_code='EG';
DELETE FROM countries WHERE iso_code='EH';
DELETE FROM countries WHERE iso_code='ER';
DELETE FROM countries WHERE iso_code='ES';
DELETE FROM countries WHERE iso_code='ET';
DELETE FROM countries WHERE iso_code='FI';
DELETE FROM countries WHERE iso_code='FJ';
DELETE FROM countries WHERE iso_code='FK';
DELETE FROM countries WHERE iso_code='FM';
DELETE FROM countries WHERE iso_code='FO';
DELETE FROM countries WHERE iso_code='FR';
DELETE FROM countries WHERE iso_code='GA';
DELETE FROM countries WHERE iso_code='GB';
DELETE FROM countries WHERE iso_code='GD';
DELETE FROM countries WHERE iso_code='GE';
DELETE FROM countries WHERE iso_code='GF';
DELETE FROM countries WHERE iso_code='GG';
DELETE FROM countries WHERE iso_code='GH';
DELETE FROM countries WHERE iso_code='GI';
DELETE FROM countries WHERE iso_code='GL';
DELETE FROM countries WHERE iso_code='GM';
DELETE FROM countries WHERE iso_code='GN';
DELETE FROM countries WHERE iso_code='GP';
DELETE FROM countries WHERE iso_code='GQ';
DELETE FROM countries WHERE iso_code='GR';
DELETE FROM countries WHERE iso_code='GT';
DELETE FROM countries WHERE iso_code='GU';
DELETE FROM countries WHERE iso_code='GW';
DELETE FROM countries WHERE iso_code='GY';
DELETE FROM countries WHERE iso_code='HK';
DELETE FROM countries WHERE iso_code='HN';
DELETE FROM countries WHERE iso_code='HR';
DELETE FROM countries WHERE iso_code='HT';
DELETE FROM countries WHERE iso_code='HU';
DELETE FROM countries WHERE iso_code='ID';
DELETE FROM countries WHERE iso_code='IE';
DELETE FROM countries WHERE iso_code='IL';
DELETE FROM countries WHERE iso_code='IM';
DELETE FROM countries WHERE iso_code='IN';
DELETE FROM countries WHERE iso_code='IO';
DELETE FROM countries WHERE iso_code='IQ';
DELETE FROM countries WHERE iso_code='IR';
DELETE FROM countries WHERE iso_code='IS';
DELETE FROM countries WHERE iso_code='IT';
DELETE FROM countries WHERE iso_code='JE';
DELETE FROM countries WHERE iso_code='JM';
DELETE FROM countries WHERE iso_code='JO';
DELETE FROM countries WHERE iso_code='JP';
DELETE FROM countries WHERE iso_code='KE';
DELETE FROM countries WHERE iso_code='KG';
DELETE FROM countries WHERE iso_code='KH';
DELETE FROM countries WHERE iso_code='KI';
DELETE FROM countries WHERE iso_code='KM';
DELETE FROM countries WHERE iso_code='KN';
DELETE FROM countries WHERE iso_code='KP';
DELETE FROM countries WHERE iso_code='KR';
DELETE FROM countries WHERE iso_code='KW';
DELETE FROM countries WHERE iso_code='KY';
DELETE FROM countries WHERE iso_code='KZ';
DELETE FROM countries WHERE iso_code='LA';
DELETE FROM countries WHERE iso_code='LB';
DELETE FROM countries WHERE iso_code='LC';
DELETE FROM countries WHERE iso_code='LI';
DELETE FROM countries WHERE iso_code='LK';
DELETE FROM countries WHERE iso_code='LR';
DELETE FROM countries WHERE iso_code='LS';
DELETE FROM countries WHERE iso_code='LT';
DELETE FROM countries WHERE iso_code='LU';
DELETE FROM countries WHERE iso_code='LV';
DELETE FROM countries WHERE iso_code='LY';
DELETE FROM countries WHERE iso_code='MA';
DELETE FROM countries WHERE iso_code='MC';
DELETE FROM countries WHERE iso_code='MD';
DELETE FROM countries WHERE iso_code='ME';
DELETE FROM countries WHERE iso_code='MF';
DELETE FROM countries WHERE iso_code='MG';
DELETE FROM countries WHERE iso_code='MH';
DELETE FROM countries WHERE iso_code='MK';
DELETE FROM countries WHERE iso_code='ML';
DELETE FROM countries WHERE iso_code='MM';
DELETE FROM countries WHERE iso_code='MN';
DELETE FROM countries WHERE iso_code='MO';
DELETE FROM countries WHERE iso_code='MP';
DELETE FROM countries WHERE iso_code='MQ';
DELETE FROM countries WHERE iso_code='MR';
DELETE FROM countries WHERE iso_code='MS';
DELETE FROM countries WHERE iso_code='MT';
DELETE FROM countries WHERE iso_code='MU';
DELETE FROM countries WHERE iso_code='MV';
DELETE FROM countries WHERE iso_code='MW';
DELETE FROM countries WHERE iso_code='MX';
DELETE FROM countries WHERE iso_code='MY';
DELETE FROM countries WHERE iso_code='MZ';
DELETE FROM countries WHERE iso_code='NA';
DELETE FROM countries WHERE iso_code='NC';
DELETE FROM countries WHERE iso_code='NE';
DELETE FROM countries WHERE iso_code='NF';
DELETE FROM countries WHERE iso_code='NG';
DELETE FROM countries WHERE iso_code='NI';
DELETE FROM countries WHERE iso_code='NL';
DELETE FROM countries WHERE iso_code='NO';
DELETE FROM countries WHERE iso_code='NP';
DELETE FROM countries WHERE iso_code='NR';
DELETE FROM countries WHERE iso_code='NU';
DELETE FROM countries WHERE iso_code='NZ';
DELETE FROM countries WHERE iso_code='OM';
DELETE FROM countries WHERE iso_code='PA';
DELETE FROM countries WHERE iso_code='PE';
DELETE FROM countries WHERE iso_code='PF';
DELETE FROM countries WHERE iso_code='PG';
DELETE FROM countries WHERE iso_code='PH';
DELETE FROM countries WHERE iso_code='PK';
DELETE FROM countries WHERE iso_code='PL';
DELETE FROM countries WHERE iso_code='PM';
DELETE FROM countries WHERE iso_code='PR';
DELETE FROM countries WHERE iso_code='PS';
DELETE FROM countries WHERE iso_code='PT';
DELETE FROM countries WHERE iso_code='PW';
DELETE FROM countries WHERE iso_code='PY';
DELETE FROM countries WHERE iso_code='QA';
DELETE FROM countries WHERE iso_code='RE';
DELETE FROM countries WHERE iso_code='RO';
DELETE FROM countries WHERE iso_code='RS';
DELETE FROM countries WHERE iso_code='RU';
DELETE FROM countries WHERE iso_code='RW';
DELETE FROM countries WHERE iso_code='SA';
DELETE FROM countries WHERE iso_code='SB';
DELETE FROM countries WHERE iso_code='SC';
DELETE FROM countries WHERE iso_code='SD';
DELETE FROM countries WHERE iso_code='SE';
DELETE FROM countries WHERE iso_code='SG';
DELETE FROM countries WHERE iso_code='SH';
DELETE FROM countries WHERE iso_code='SI';
DELETE FROM countries WHERE iso_code='SJ';
DELETE FROM countries WHERE iso_code='SK';
DELETE FROM countries WHERE iso_code='SL';
DELETE FROM countries WHERE iso_code='SM';
DELETE FROM countries WHERE iso_code='SN';
DELETE FROM countries WHERE iso_code='SO';
DELETE FROM countries WHERE iso_code='SR';
DELETE FROM countries WHERE iso_code='SS';
DELETE FROM countries WHERE iso_code='ST';
DELETE FROM countries WHERE iso_code='SV';
DELETE FROM countries WHERE iso_code='SX';
DELETE FROM countries WHERE iso_code='SY';
DELETE FROM countries WHERE iso_code='SZ';
DELETE FROM countries WHERE iso_code='TA';
DELETE FROM countries WHERE iso_code='TC';
DELETE FROM countries WHERE iso_code='TD';
DELETE FROM countries WHERE iso_code='TG';
DELETE FROM countries WHERE iso_code='TH';
DELETE FROM countries WHERE iso_code='TJ';
DELETE FROM countries WHERE iso_code='TK';
DELETE FROM countries WHERE iso_code='TL';
DELETE FROM countries WHERE iso_code='TM';
DELETE FROM countries WHERE iso_code='TN';
DELETE FROM countries WHERE iso_code='TO';
DELETE FROM countries WHERE iso_code='TR';
DELETE FROM countries WHERE iso_code='TT';
DELETE FROM countries WHERE iso_code='TV';
DELETE FROM countries WHERE iso_code='TW';
DELETE FROM countries WHERE iso_code='TZ';
DELETE FROM countries WHERE iso_code='UA';
DELETE FROM countries WHERE iso_code='UG';
DELETE FROM countries WHERE iso_code='US';
DELETE FROM countries WHERE iso_code='UY';
DELETE FROM countries WHERE iso_code='UZ';
DELETE FROM countries WHERE iso_code='VA';
DELETE FROM countries WHERE iso_code='VC';
DELETE FROM countries WHERE iso_code='VE';
DELETE FROM countries WHERE iso_code='VG';
DELETE FROM countries WHERE iso_code='VI';
DELETE FROM countries WHERE iso_code='VN';
DELETE FROM countries WHERE iso_code='VU';
DELETE FROM countries WHERE iso_code='WF';
DELETE FROM countries WHERE iso_code='WS';
DELETE FROM countries WHERE iso_code='XK';
DELETE FROM countries WHERE iso_code='YE';
DELETE FROM countries WHERE iso_code='YT';
DELETE FROM countries WHERE iso_code='ZA';
DELETE FROM countries WHERE iso_code='ZM';
DELETE FROM countries WHERE iso_code='ZW';
-- +goose StatementEnd
