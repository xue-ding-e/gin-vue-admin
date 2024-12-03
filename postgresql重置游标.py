import psycopg2
# pip install psycopg2 -i https://pypi.org/simple
# pip install -r requirements.txt -i https://pypi.org/simple
def reset_sequences():
    conn_params = {
        'host': '',
        'port': 5432,
        'database': '',
        'user': '',
        'password': ''
    }

    conn = psycopg2.connect(**conn_params)
    conn.autocommit = True
    cur = conn.cursor()

    # 获取所有主键列
    cur.execute("""
    SELECT
        kcu.table_schema,
        kcu.table_name,
        kcu.column_name
    FROM
        information_schema.key_column_usage AS kcu
    JOIN
        information_schema.table_constraints AS tc
    ON
        kcu.constraint_name = tc.constraint_name
        AND kcu.table_schema = tc.table_schema
    WHERE
        tc.constraint_type = 'PRIMARY KEY'
        AND kcu.table_schema NOT IN ('pg_catalog', 'information_schema');
    """)

    pk_columns = cur.fetchall()

    for schema_name, table_name, column_name in pk_columns:
        seq_name = '{}_{}_seq'.format(table_name, column_name)
        full_seq_name = '{}.{}'.format(schema_name, seq_name)
        full_table_name = '{}.{}'.format(schema_name, table_name)

        # 检查序列是否存在
        cur.execute("""
        SELECT EXISTS (
            SELECT 1 FROM pg_class WHERE relkind = 'S' AND relname = %s
        );
        """, (seq_name,))
        seq_exists = cur.fetchone()[0]

        if not seq_exists:
            cur.execute('CREATE SEQUENCE {}.{}'.format(schema_name, seq_name))
            print('序列 {} 已创建'.format(full_seq_name))
        else:
            print('序列 {} 已存在'.format(full_seq_name))

        # 检查列的默认值
        cur.execute("""
        SELECT column_default FROM information_schema.columns
        WHERE table_schema = %s AND table_name = %s AND column_name = %s;
        """, (schema_name, table_name, column_name))
        col_default = cur.fetchone()[0]

        if not col_default or 'nextval' not in col_default:
            cur.execute("""
            ALTER TABLE {}.{} ALTER COLUMN {} SET DEFAULT nextval('{}'::regclass);
            """.format(schema_name, table_name, column_name, full_seq_name))
            print('表 {}.{} 的列 {} 的默认值已设置为 nextval({})'.format(schema_name, table_name, column_name, full_seq_name))
        else:
            print('表 {}.{} 的列 {} 已有默认值'.format(schema_name, table_name, column_name))

        # 设置序列的所有权
        cur.execute("""
        ALTER SEQUENCE {}.{} OWNED BY {}.{}.{};
        """.format(schema_name, seq_name, schema_name, table_name, column_name))

        # 获取主键列的最大值
        cur.execute('SELECT MAX({}) FROM {}.{}'.format(column_name, schema_name, table_name))
        max_id = cur.fetchone()[0]

        if max_id is None:
            max_id = 1
        else:
            max_id += 1

        # 重置序列的值
        cur.execute("SELECT setval('{}', {}, false);".format(full_seq_name, max_id))
        print('序列 {} 已重置为 {}'.format(full_seq_name, max_id))

    cur.close()
    conn.close()
    print('序列重置完成！')

if __name__ == '__main__':
    reset_sequences()