import { Type as T } from '@sinclair/typebox'
import {
  UserDTO,
  UserTypeSchema,
} from '../../../../../../domain/entities/user/user'
import { Email } from '../../../../../../domain/models/email'
import { CreateUserCommandFactory } from '../../../../../../domain/use-cases/commands/create-user/command'
import { ServiceFastifyInstance } from '../../../../api'

export type SignupUserResponse = {
  result: UserDTO
}

export const SignupUserRoute = async (fastify: ServiceFastifyInstance) =>
  fastify.put(
    '/v1/users/signup',
    {
      schema: {
        body: T.Object({
          name: T.String(),
          email: T.String({ format: 'email' }),
          password: T.String(),
        }),
        response: {
          200: {
            result: {
              user: UserTypeSchema,
              accessToken: T.String(),
              refreshToken: T.String(),
            },
          },
        },
      },
    },
    async (request): Promise<SignupUserResponse> => {
      const command = new CreateUserCommandFactory().instance(fastify.appCtx)
      const user = await command.execute({
        name: request.body.name,
        email: new Email(request.body.email),
      })

      return {
        result: user.serialize(),
      }
    },
  )
