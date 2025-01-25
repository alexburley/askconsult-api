import * as crypto from 'crypto'
import { UserRepository } from '../../../../adapters/repositories/user'
import { User, UserTypes } from '../../../entities/user/user'
import { Email } from '../../../models/email'
import { ApplicationContext } from '../../../../lib/app-ctx/app-ctx'
import { UserRepositoryFactory } from '../../../../adapters/repositories/user/factory'

export class SignupUserCommandFactory {
  instance(ctx: ApplicationContext) {
    return new SignupUserCommand(ctx, {
      users: new UserRepositoryFactory().instance(ctx),
    })
  }
}

export type CreateUserCommandDeps = {
  users: UserRepository
}

export class SignupUserCommand {
  users

  constructor(ctx: ApplicationContext, deps: CreateUserCommandDeps) {
    this.users = deps.users
  }

  async execute(input: { name: string; email: Email; password: string }) {
    const existingUser = await this.users.getByEmail(input.email)

    if (existingUser) {
      throw new Error('User already exists')
    }

    const salt = crypto.randomBytes(16).toString('hex')
    const hash = crypto
      .pbkdf2Sync(input.password, salt, 1000, 64, 'sha512')
      .toString('hex')

    const user = new User({
      name: input.name,
      email: input.email,
      password: {
        hash,
        salt,
      },
      userType: UserTypes.Customer,
    })

    await this.users.persist(user)

    return user
  }
}
