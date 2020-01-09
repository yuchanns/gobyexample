## struct of tables

```sql
CREATE TABLE `articles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `mid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'markdown_id',
  `cid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'category_id',
  `title` varchar(20) NOT NULL DEFAULT '',
  `excerpt` varchar(50) NOT NULL DEFAULT '',
  `content` mediumtext,
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
  `deleted_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `markdown` (`mid`) USING BTREE,
  KEY `category` (`cid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='articles';
```

```sql
CREATE TABLE `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL DEFAULT '',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='categories';
```

```sql
CREATE TABLE `markdown` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `content` mediumtext,
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `is_draft` tinyint(1) unsigned NOT NULL DEFAULT '0',
  `created_at` int(10) unsigned NOT NULL DEFAULT '0',
  `updated_at` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `draft` (`is_draft`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='original markdown';
```

```sql
CREATE TABLE `tag_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `tid` int(10) unsigned NOT NULL DEFAULT '0',
  `aid` int(10) unsigned NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `article` (`aid`),
  KEY `tag` (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='m2m of tags and articles';
```

```sql
CREATE TABLE `tags` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(5) NOT NULL DEFAULT '',
  `is_deleted` tinyint(1) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='tags';
```

