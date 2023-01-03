// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/babyname/fate/ent/migrate"

	"github.com/babyname/fate/ent/character"
	"github.com/babyname/fate/ent/version"
	"github.com/babyname/fate/ent/wugelucky"
	"github.com/babyname/fate/ent/wuxing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Character is the client for interacting with the Character builders.
	Character *CharacterClient
	// Version is the client for interacting with the Version builders.
	Version *VersionClient
	// WuGeLucky is the client for interacting with the WuGeLucky builders.
	WuGeLucky *WuGeLuckyClient
	// WuXing is the client for interacting with the WuXing builders.
	WuXing *WuXingClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Character = NewCharacterClient(c.config)
	c.Version = NewVersionClient(c.config)
	c.WuGeLucky = NewWuGeLuckyClient(c.config)
	c.WuXing = NewWuXingClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Character: NewCharacterClient(cfg),
		Version:   NewVersionClient(cfg),
		WuGeLucky: NewWuGeLuckyClient(cfg),
		WuXing:    NewWuXingClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Character: NewCharacterClient(cfg),
		Version:   NewVersionClient(cfg),
		WuGeLucky: NewWuGeLuckyClient(cfg),
		WuXing:    NewWuXingClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Character.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Character.Use(hooks...)
	c.Version.Use(hooks...)
	c.WuGeLucky.Use(hooks...)
	c.WuXing.Use(hooks...)
}

// CharacterClient is a client for the Character schema.
type CharacterClient struct {
	config
}

// NewCharacterClient returns a client for the Character from the given config.
func NewCharacterClient(c config) *CharacterClient {
	return &CharacterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `character.Hooks(f(g(h())))`.
func (c *CharacterClient) Use(hooks ...Hook) {
	c.hooks.Character = append(c.hooks.Character, hooks...)
}

// Create returns a builder for creating a Character entity.
func (c *CharacterClient) Create() *CharacterCreate {
	mutation := newCharacterMutation(c.config, OpCreate)
	return &CharacterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Character entities.
func (c *CharacterClient) CreateBulk(builders ...*CharacterCreate) *CharacterCreateBulk {
	return &CharacterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Character.
func (c *CharacterClient) Update() *CharacterUpdate {
	mutation := newCharacterMutation(c.config, OpUpdate)
	return &CharacterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CharacterClient) UpdateOne(ch *Character) *CharacterUpdateOne {
	mutation := newCharacterMutation(c.config, OpUpdateOne, withCharacter(ch))
	return &CharacterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CharacterClient) UpdateOneID(id string) *CharacterUpdateOne {
	mutation := newCharacterMutation(c.config, OpUpdateOne, withCharacterID(id))
	return &CharacterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Character.
func (c *CharacterClient) Delete() *CharacterDelete {
	mutation := newCharacterMutation(c.config, OpDelete)
	return &CharacterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CharacterClient) DeleteOne(ch *Character) *CharacterDeleteOne {
	return c.DeleteOneID(ch.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *CharacterClient) DeleteOneID(id string) *CharacterDeleteOne {
	builder := c.Delete().Where(character.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CharacterDeleteOne{builder}
}

// Query returns a query builder for Character.
func (c *CharacterClient) Query() *CharacterQuery {
	return &CharacterQuery{
		config: c.config,
	}
}

// Get returns a Character entity by its id.
func (c *CharacterClient) Get(ctx context.Context, id string) (*Character, error) {
	return c.Query().Where(character.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CharacterClient) GetX(ctx context.Context, id string) *Character {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CharacterClient) Hooks() []Hook {
	return c.hooks.Character
}

// VersionClient is a client for the Version schema.
type VersionClient struct {
	config
}

// NewVersionClient returns a client for the Version from the given config.
func NewVersionClient(c config) *VersionClient {
	return &VersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `version.Hooks(f(g(h())))`.
func (c *VersionClient) Use(hooks ...Hook) {
	c.hooks.Version = append(c.hooks.Version, hooks...)
}

// Create returns a builder for creating a Version entity.
func (c *VersionClient) Create() *VersionCreate {
	mutation := newVersionMutation(c.config, OpCreate)
	return &VersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Version entities.
func (c *VersionClient) CreateBulk(builders ...*VersionCreate) *VersionCreateBulk {
	return &VersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Version.
func (c *VersionClient) Update() *VersionUpdate {
	mutation := newVersionMutation(c.config, OpUpdate)
	return &VersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VersionClient) UpdateOne(v *Version) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersion(v))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VersionClient) UpdateOneID(id int) *VersionUpdateOne {
	mutation := newVersionMutation(c.config, OpUpdateOne, withVersionID(id))
	return &VersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Version.
func (c *VersionClient) Delete() *VersionDelete {
	mutation := newVersionMutation(c.config, OpDelete)
	return &VersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *VersionClient) DeleteOne(v *Version) *VersionDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *VersionClient) DeleteOneID(id int) *VersionDeleteOne {
	builder := c.Delete().Where(version.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VersionDeleteOne{builder}
}

// Query returns a query builder for Version.
func (c *VersionClient) Query() *VersionQuery {
	return &VersionQuery{
		config: c.config,
	}
}

// Get returns a Version entity by its id.
func (c *VersionClient) Get(ctx context.Context, id int) (*Version, error) {
	return c.Query().Where(version.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VersionClient) GetX(ctx context.Context, id int) *Version {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *VersionClient) Hooks() []Hook {
	return c.hooks.Version
}

// WuGeLuckyClient is a client for the WuGeLucky schema.
type WuGeLuckyClient struct {
	config
}

// NewWuGeLuckyClient returns a client for the WuGeLucky from the given config.
func NewWuGeLuckyClient(c config) *WuGeLuckyClient {
	return &WuGeLuckyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `wugelucky.Hooks(f(g(h())))`.
func (c *WuGeLuckyClient) Use(hooks ...Hook) {
	c.hooks.WuGeLucky = append(c.hooks.WuGeLucky, hooks...)
}

// Create returns a builder for creating a WuGeLucky entity.
func (c *WuGeLuckyClient) Create() *WuGeLuckyCreate {
	mutation := newWuGeLuckyMutation(c.config, OpCreate)
	return &WuGeLuckyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WuGeLucky entities.
func (c *WuGeLuckyClient) CreateBulk(builders ...*WuGeLuckyCreate) *WuGeLuckyCreateBulk {
	return &WuGeLuckyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WuGeLucky.
func (c *WuGeLuckyClient) Update() *WuGeLuckyUpdate {
	mutation := newWuGeLuckyMutation(c.config, OpUpdate)
	return &WuGeLuckyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WuGeLuckyClient) UpdateOne(wgl *WuGeLucky) *WuGeLuckyUpdateOne {
	mutation := newWuGeLuckyMutation(c.config, OpUpdateOne, withWuGeLucky(wgl))
	return &WuGeLuckyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WuGeLuckyClient) UpdateOneID(id string) *WuGeLuckyUpdateOne {
	mutation := newWuGeLuckyMutation(c.config, OpUpdateOne, withWuGeLuckyID(id))
	return &WuGeLuckyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WuGeLucky.
func (c *WuGeLuckyClient) Delete() *WuGeLuckyDelete {
	mutation := newWuGeLuckyMutation(c.config, OpDelete)
	return &WuGeLuckyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WuGeLuckyClient) DeleteOne(wgl *WuGeLucky) *WuGeLuckyDeleteOne {
	return c.DeleteOneID(wgl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WuGeLuckyClient) DeleteOneID(id string) *WuGeLuckyDeleteOne {
	builder := c.Delete().Where(wugelucky.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WuGeLuckyDeleteOne{builder}
}

// Query returns a query builder for WuGeLucky.
func (c *WuGeLuckyClient) Query() *WuGeLuckyQuery {
	return &WuGeLuckyQuery{
		config: c.config,
	}
}

// Get returns a WuGeLucky entity by its id.
func (c *WuGeLuckyClient) Get(ctx context.Context, id string) (*WuGeLucky, error) {
	return c.Query().Where(wugelucky.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WuGeLuckyClient) GetX(ctx context.Context, id string) *WuGeLucky {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WuGeLuckyClient) Hooks() []Hook {
	return c.hooks.WuGeLucky
}

// WuXingClient is a client for the WuXing schema.
type WuXingClient struct {
	config
}

// NewWuXingClient returns a client for the WuXing from the given config.
func NewWuXingClient(c config) *WuXingClient {
	return &WuXingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `wuxing.Hooks(f(g(h())))`.
func (c *WuXingClient) Use(hooks ...Hook) {
	c.hooks.WuXing = append(c.hooks.WuXing, hooks...)
}

// Create returns a builder for creating a WuXing entity.
func (c *WuXingClient) Create() *WuXingCreate {
	mutation := newWuXingMutation(c.config, OpCreate)
	return &WuXingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WuXing entities.
func (c *WuXingClient) CreateBulk(builders ...*WuXingCreate) *WuXingCreateBulk {
	return &WuXingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WuXing.
func (c *WuXingClient) Update() *WuXingUpdate {
	mutation := newWuXingMutation(c.config, OpUpdate)
	return &WuXingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WuXingClient) UpdateOne(wx *WuXing) *WuXingUpdateOne {
	mutation := newWuXingMutation(c.config, OpUpdateOne, withWuXing(wx))
	return &WuXingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WuXingClient) UpdateOneID(id string) *WuXingUpdateOne {
	mutation := newWuXingMutation(c.config, OpUpdateOne, withWuXingID(id))
	return &WuXingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WuXing.
func (c *WuXingClient) Delete() *WuXingDelete {
	mutation := newWuXingMutation(c.config, OpDelete)
	return &WuXingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WuXingClient) DeleteOne(wx *WuXing) *WuXingDeleteOne {
	return c.DeleteOneID(wx.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WuXingClient) DeleteOneID(id string) *WuXingDeleteOne {
	builder := c.Delete().Where(wuxing.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WuXingDeleteOne{builder}
}

// Query returns a query builder for WuXing.
func (c *WuXingClient) Query() *WuXingQuery {
	return &WuXingQuery{
		config: c.config,
	}
}

// Get returns a WuXing entity by its id.
func (c *WuXingClient) Get(ctx context.Context, id string) (*WuXing, error) {
	return c.Query().Where(wuxing.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WuXingClient) GetX(ctx context.Context, id string) *WuXing {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *WuXingClient) Hooks() []Hook {
	return c.hooks.WuXing
}
