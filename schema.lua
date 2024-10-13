box.cfg{}

box.schema.space.create('users', {
    format = {
        {name = 'id', type = 'string'},
        {name = 'balance', type = 'number'}
    }
})

box.space.users:create_index('primary', {
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

box.space.rates:insert{0.000016,'BTC',os.time(os.date("*t"))}
box.space.rates:insert{0.00041006,'ETH',os.time(os.date("*t"))}

box.schema.space.create('orders', {
    format = {
        {name = 'id', type = 'string'},
        {name = 'user_id', type = 'string'},
        {name = 'market', type = 'string'},
        {name = 'side', type = 'string'},
        {name = 'price', type = 'number'},
        {name = 'size', type = 'number'},
        {name = 'status', type = 'string'}
    }
})

box.space.orders:create_index('primary', {
    type = 'tree',
    parts = {'id'},
    unique = true
})

box.space.orders:create_index('user_id', {
    type = 'tree',
    parts = {'user_id'},
    unique = false
})

box.space.orders:create_index('side', {
    type = 'tree',
    parts = {'side'},
    unique = false
})

box.space.orders:create_index('todo_orders', {
    type = 'tree',
    parts = {'side', 'status'},
    unique = false
})

box.schema.space.create('positions', {
    format = {
        {name = 'user_id', type = 'string'},
        {name = 'market', type = 'string'},
        {name = 'side', type = 'string'},
        {name = 'entry_price', type = 'number'},
        {name = 'size', type = 'number'}
    }
})

box.space.positions:create_index('primary', {
    type = 'tree',
    parts = {'user_id', 'market', 'side'},
    unique = true
})

box.space.positions:create_index('user_id', {
    type = 'tree',
    parts = {'user_id'},
    unique = false
})
