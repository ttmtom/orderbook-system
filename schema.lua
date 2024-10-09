box.cfg{}

box.schema.space.create('users', {
    format = {
        {name = 'id', type = 'unsigned'},
    }
})

box.space.users:create_index('primary', {
    type = 'tree',
    parts = {'id'},
    unique = true
})

box.schema.space.create('accounts', {
    format = {
        {name = 'user_id', type = 'unsigned'},
        {name = 'amount', type = 'number'},
        {name = 'currency', type = 'string'}
    }
})

box.space.accounts:create_index('primary', {
    type = 'tree',
    parts = {'user_id', 'currency'},
    unique = true
})

box.schema.space.create('orders', {
    format = {
        {name = 'id', type = 'unsigned'},
        {name = 'user_id', type = 'unsigned'},
        {name = 'rate', type = 'number'},
        {name = 'currency', type = 'string'},
        {name = 'amount', type = 'number'},
        {name = 'type', type = 'string'},
        {name = 'status', type = 'string'},
        {name = 'created_at', type = 'unsigned'},
        {name = 'updated_at', type = 'unsigned'},
        {name = 'pair', type = 'unsigned'}
    }
})

box.space.orders:create_index('primary', {
    type = 'tree',
    parts = {'id'},
    unique = true
})

box.schema.space.create('rates', {
    format = {
        {name = 'rate', type = 'number'},
        {name = 'currency', type = 'string'},
        {name = 'updated_at', type = 'unsigned'}
    }
})

box.space.rates:create_index('primary', {
    type = 'tree',
    parts = {'currency'},
    unique = true
})
