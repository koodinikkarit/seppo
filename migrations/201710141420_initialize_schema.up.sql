
create table if not exists tags (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) DEFAULT "" NOT NULL,
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL
);

create table if not exists languages (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) DEFAULT "" NOT NULL,
	created_at DATETIME,
	updated_at DATETIME NULL
);

create table if not exists logs (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	log_type INT NOT NULL,
	message TEXT NOT NULL,
	message_date DATETIME
);

create table if not exists songs(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY
);

create table if not exists schedules(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) DEFAULT "" NOT NULL,
	start DATETIME NULL,
	end DATETIME NULL,
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL
);

create table if not exists events(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50),
	start DATETIME,
	end DATETIME,
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL	
);

create table if not exists authors(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255) DEFAULT "" NOT NULL,
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL
);

create table if not exists copyrights(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(255) DEFAULT "" NOT NULL,
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL
);

create table if not exists variations (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	song_id INT8 UNSIGNED DEFAULT NULL,
	language_id INT8 UNSIGNED DEFAULT NULL,
	variation_id INT8 UNSIGNED DEFAULT NULL,
	ew_song_id INT8 UNSIGNED DEFAULT NULL,
	jyvaskyla_song_id INT8 UNSIGNED DEFAULT NULL,
	author_id INT8 UNSIGNED NULL,
	copyright_id INT8 UNSIGNED NULL,
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME DEFAULT NULL,
	FOREIGN KEY(song_id) REFERENCES songs(id),
	FOREIGN KEY(language_id) REFERENCES languages(id),
	FOREIGN KEY(author_id) REFERENCES authors(id),
	FOREIGN KEY(copyright_id) REFERENCES copyrights(id)
);

create table if not exists jyvaskyla_songs(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	added_at INT8 UNSIGNED,
	added_by VARCHAR(50),
	additional_info VARCHAR(255),
	arrangement_by VARCHAR(50),
	composed_by VARCHAR(50),
	copyright VARCHAR(50),
	deleted BOOLEAN,
	global_id INT8 UNSIGNED,
	lyrics_by VARCHAR(50),
	modified DATETIME,
	name VARCHAR(50),
	orig_name VARCHAR(50),
	song TEXT,
	songbook_id INT8 UNSIGNED,
	translated_by VARCHAR(50),
	year VARCHAR(50)
);

create table if not exists ew_songs(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	title VARCHAR(50),
	author VARCHAR(50),
	copyright VARCHAR(50),
	administrator VARCHAR(50),
	description VARCHAR(50),
	tags VARCHAR(50),
	text TEXT
);

create table if not exists variation_versions (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	variation_id INT8 UNSIGNED NOT NULL,
	name VARCHAR(255) NOT NULL,
	text TEXT NOT NULL,
	version INT UNSIGNED DEFAULT 1 NOT NULL,
	newest BOOLEAN,
	created_at DATETIME,
	disabled_at DATETIME NULL,
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);

create table if not exists branches (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	source_variation_version_id INT8 UNSIGNED NOT NULL,
	destination_variation_version_id INT8 UNSIGNED NOT NULL,
	created_at DATETIME,
	FOREIGN KEY(source_variation_version_id) REFERENCES variation_versions(id),
	FOREIGN KEY(destination_variation_version_id) REFERENCES variation_versions(id)
);

create table if not exists merges (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	variation_version1_id INT8 UNSIGNED NOT NULL,
	variation_version2_id INT8 UNSIGNED NOT NULL,
	destination_variation_version_id INT8 UNSIGNED NOT NULL,
	created_at DATETIME,
	FOREIGN KEY(variation_version1_id) REFERENCES variation_versions(id),
	FOREIGN KEY(variation_version2_id) REFERENCES variation_versions(id),
	FOREIGN KEY(destination_variation_version_id) REFERENCES variation_versions(id)
);

create table if not exists song_databases (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) DEFAULT "" NOT NULL,
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL
);

create table if not exists matias_client (
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	client_key VARCHAR(20),
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL	
);

create table if not exists ew_databases(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	name VARCHAR(50) DEFAULT "" NOT NULL,
	song_database_id INT8 UNSIGNED NOT NULL,
	ew_database_key VARCHAR(10) DEFAULT "" NOT NULL,
	use_newest_version BOOLEAN,
	matias_client_id INT8 UNSIGNED NULL,
	remove_songs_from_ew_database BOOLEAN DEFAULT false NOT NULL,
	remove_songs_from_song_database BOOLEAN DEFAULT false NOT NULL,
	variation_version_conflict_action INT UNSIGNED DEFAULT 0 NOT NULL, 
	created_at DATETIME,
	updated_at DATETIME NULL,
	deleted_at DATETIME NULL,
	FOREIGN KEY(song_database_id) REFERENCES song_databases(id),
	FOREIGN KEY(matias_client_id) REFERENCES matias_client(id)
);

create table if not exists ew_database_links(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	ew_database_id INT8 UNSIGNED NOT NULL,
	ew_database_song_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	version INT UNSIGNED DEFAULT 1 NOT NULL,
	author VARCHAR(255) DEFAULT "" NOT NULL,
	copyright VARCHAR(255) DEFAULT "" NOT NULL,
	created_at DATETIME,
	updated_at DATETIME,
	FOREIGN KEY(ew_database_id) REFERENCES ew_databases(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);

create table if not exists song_database_variations(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	song_database_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	created_at DATETIME,
	deleted_at DATETIME NULL,
	FOREIGN KEY(song_database_id) REFERENCES song_databases(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);

create table if not exists song_database_tags(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	tag_id INT8 UNSIGNED NOT NULL,
	song_database_id INT8 UNSIGNED NOT NULL,
	created_at DATETIME,
	deleted_at DATETIME NULL,
	FOREIGN KEY(tag_id) REFERENCES tags(id),
	FOREIGN KEY(song_database_id) REFERENCES song_databases(id)
);

create table if not exists tag_variations(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	tag_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	created_at DATETIME,
	FOREIGN KEY(tag_id) REFERENCES tags(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);

create table if not exists schedule_variations(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	schedule_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	order_number INT,
	created_at DATETIME,
	deleted_at DATETIME NULL,
	FOREIGN KEY(schedule_id) REFERENCES schedules(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);

create table if not exists event_schedules(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	schedule_id INT8 UNSIGNED NOT NULL,
	event_id INT8 UNSIGNED NOT NULL,
	created_at DATETIME,
	deleted_at DATETIME NULL,
	FOREIGN KEY(schedule_id) REFERENCES schedules(id),
	FOREIGN KEY(event_id) REFERENCES events(id)
);

create table if not exists synchronization_raports(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	raport_type INT,
	database_id INT8 UNSIGNED NULL,
	database_key VARCHAR(10),
	database_found BOOLEAN,
	duration_ms INT8,
	sr_gen_duration_ms INT8,
	started_at DATETIME,
	finished_at DATETIME
);

create table if not exists sr_ew_conflicts(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	variation_version_id INT8 UNSIGNED NOT NULL,
	ew_database_id INT8 UNSIGNED NOT NULL,
	ew_song_id INT8 UNSIGNED NOT NULL,
	name VARCHAR(50),
	text TEXT,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
 	FOREIGN KEY(variation_version_id) REFERENCES variation_versions(id),
	FOREIGN KEY(ew_database_id) REFERENCES ew_databases(id)
);

create table if not exists sr_new_authors(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	author_id INT8 UNSIGNED NOT NULL,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(author_id) REFERENCES authors(id)
);

create table if not exists sr_new_copyrights(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	copyright_id INT8 UNSIGNED NOT NULL,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(copyright_id) REFERENCES copyrights(id)
);

create table if not exists sr_new_variations(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);

create table if not exists sr_passivated_variation_versions(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	variation_version_id INT8 UNSIGNED NOT NULL,	
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(variation_version_id) REFERENCES variation_versions(id)
);

create table if not exists sr_ew_database_links(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	ew_database_id INT8 UNSIGNED NOT NULL,
	ew_database_song_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	version INT UNSIGNED DEFAULT 1 NOT NULL,
	author VARCHAR(255) DEFAULT "" NOT NULL,
	copyright VARCHAR(255) DEFAULT "" NOT NULL,
	operation BOOLEAN NOT NULL,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(ew_database_id) REFERENCES ew_databases(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);

create table if not exists sr_new_branches(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	branch_id INT8 UNSIGNED NOT NULL,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(branch_id) REFERENCES branches(id)
);

create table if not exists sr_new_variation_versions(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	variation_version_id INT8 UNSIGNED NOT NULL,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(variation_version_id) REFERENCES variation_versions(id)
);

create table if not exists sr_ew_song(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,	
	variation_version_id INT8 UNSIGNED NOT NULL,
	operation BOOLEAN,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(variation_version_id) REFERENCES variation_versions(id)
);

create table if not exists sr_add_song_database_variations(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	song_database_id INT8 UNSIGNED NOT NULL,	
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id),
	FOREIGN KEY(song_database_id) REFERENCES song_databases(id)
);

create table if not exists sr_remove_song_database_variations(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	song_database_id INT8 UNSIGNED NOT NULL,	
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id),
	FOREIGN KEY(song_database_id) REFERENCES song_databases(id)
);

create table if not exists sr_updated_ew_database_link_versions(
	id INT8 UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	sr_id INT8 UNSIGNED NOT NULL,
	orig_version INT UNSIGNED NOT NULL,
	changed_version INT UNSIGNED NOT NULL,
	variation_id INT8 UNSIGNED NOT NULL,
	FOREIGN KEY(sr_id) REFERENCES synchronization_raports(id),
	FOREIGN KEY(variation_id) REFERENCES variations(id)
);