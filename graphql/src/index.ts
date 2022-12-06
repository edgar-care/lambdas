import { ApolloServer } from '@apollo/server';
import { startServerAndCreateLambdaHandler } from '@as-integrations/aws-lambda';
import { startStandaloneServer } from '@apollo/server/standalone';
import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient();

const typeDefs = `#graphql

    type User {
        id:  ID!
        email: String!
        name: String
        age: Int
    }

    type Query {
        getUsers: [User]
        findUser(id: String): User
    }

    type Mutation {
        createUser(email: String!, name: String, age: Int): User
        updateUser(id: ID!, email: String, name: String, age: Int): User
        deleteUser(id: ID!): Boolean
    }
`;

type UserCreateInput = { email: string, name?: string, age?: number };
type UserUpdateInput = { id: string, email?: string, name?: string, age?: number };
type UserDeleteInput = { id: string };

const resolvers = {
    Query: {
        getUsers: async () => await prisma.user.findMany(),
        findUser: async (id: string) => await prisma.user.findUnique({ where: { id }})
    },

    Mutation: {
        createUser: async (_: any, { email, name, age }: UserCreateInput) => await prisma.user.create({
                data: {
                    email,
                    name : name || null,
                    age: age ||null
                }
            }),
        updateUser: async (_ :any, { id, email, name, age }: UserUpdateInput) => await prisma.user.update({ where: { id }, data: { email, name, age } }),
        deleteUser: async (_: any, { id }: UserDeleteInput) => {
            try {
                await prisma.user.delete({ where: { id } })
                return true
            }
            catch(e) {
                return false
            }
        }
    }
};

const server = new ApolloServer({
    typeDefs,
    resolvers,
});

// Uncomment while local debugging
const { url } = await startStandaloneServer(server)
console.log("Server running at " + url)

// Comment while local debugging
// export const graphqlHandler = startServerAndCreateLambdaHandler(server);
