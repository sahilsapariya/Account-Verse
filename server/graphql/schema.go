package graphql

import (    
    "github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
    Name: "User",
    Fields: graphql.Fields{
        "id": &graphql.Field{
            Type: graphql.Int,
        },
        "name": &graphql.Field{
            Type: graphql.String,
        },
        "email": &graphql.Field{
            Type: graphql.String,
        },
        "age": &graphql.Field{
            Type: graphql.Int,
        },
    },
})

var UserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
    Name: "UserInput",
    Fields: graphql.InputObjectConfigFieldMap{
        "name": &graphql.InputObjectFieldConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "email": &graphql.InputObjectFieldConfig{
            Type: graphql.NewNonNull(graphql.String),
        },
        "age": &graphql.InputObjectFieldConfig{
            Type: graphql.Int,
        },
    },
})

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
    Name: "RootQuery",
    Fields: graphql.Fields{
        "users": &graphql.Field{
            Type:    graphql.NewList(UserType),
            Resolve: GetUsersResolver,
        },
        "user": &graphql.Field{
            Type: UserType,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.Int),
                },
            },
            Resolve: GetUserResolver,
        },
    },
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
    Name: "RootMutation",
    Fields: graphql.Fields{
        "createUser": &graphql.Field{
            Type: UserType,
            Args: graphql.FieldConfigArgument{
                "input": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(UserInputType),
                },
            },
            Resolve: CreateUserResolver,
        },
        "updateUser": &graphql.Field{
            Type: UserType,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.Int),
                },
                "input": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(UserInputType),
                },
            },
            Resolve: UpdateUserResolver,
        },
        "deleteUser": &graphql.Field{
            Type: graphql.Boolean,
            Args: graphql.FieldConfigArgument{
                "id": &graphql.ArgumentConfig{
                    Type: graphql.NewNonNull(graphql.Int),
                },
            },
            Resolve: DeleteUserResolver,
        },
    },
})

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query:    RootQuery,
    Mutation: RootMutation,
})