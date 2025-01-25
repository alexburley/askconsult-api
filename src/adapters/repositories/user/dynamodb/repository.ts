import { PaginationOptions, UserRepository } from '..'
import {
  User,
  UserStatus,
  UserType,
} from '../../../../domain/entities/user/user'
import { Email } from '../../../../domain/models/email'

import {
  cursorToStartKey,
  lastEvaluatedKeyToCursor,
} from '../../../../lib/clients/dynamodb/dynamodb'
import { UserNotFoundError } from '../errors'
import { UserSchemaDDB } from './ddb-schema'

export type UserDynamoDBRepositoryDeps = {
  ddbSchema: UserSchemaDDB
  pk_skIndex: string
}

const _UserFromDynamoDB = (Item: {
  id: string
  email: string
  name: string
  status: string
  passwordHash: string
  passwordSaltHash: string
  userType: string
  created: string
  modified: string
  deletedAt?: string
}) => {
  return new User({
    id: Item.id,
    email: new Email(Item.email),
    name: Item.name,
    status: Item.status as UserStatus,
    password: {
      hash: Item.passwordHash,
      salt: Item.passwordSaltHash
    },
    userType: Item.userType as UserType,
    createdAt: new Date(Item.created),
    modifiedAt: new Date(Item.modified),
    deletedAt: Item.deletedAt ? new Date(Item.deletedAt) : undefined,
  })
}

export class UserDynamoDBRepository implements UserRepository {
  schema
  pk_skIndex

  constructor(deps: UserDynamoDBRepositoryDeps) {
    this.schema = deps.ddbSchema
    this.pk_skIndex = deps.pk_skIndex
  }

  async getById(id: string) {
    const { Item } = await this.schema.get({ id, sk: 'USER#' })

    if (!Item || Item.deletedAt) {
      throw new UserNotFoundError(id)
    }

    return _UserFromDynamoDB(Item)
  }

  async getByEmail(): Promise<User> {
    throw new Error('Method not implemented.')
  }

  async list(opts: PaginationOptions = {}) {
    const { limit = 10, cursor: nextToken } = opts
    const { Items = [], LastEvaluatedKey } = await this.schema.query('USER#', {
      index: this.pk_skIndex,
      startKey: cursorToStartKey(nextToken),
      limit,
    })

    const collection = Items.map(_UserFromDynamoDB)

    return { collection, cursor: lastEvaluatedKeyToCursor(LastEvaluatedKey) }
  }

  async persist(user: User) {
    await this.schema.put({
      id: user.id,
      email: user.email.value,
      name: user.name,
      status: user.status,
      passwordHash: user.password.hash,
      passwordSaltHash: user.password.salt,
      userType: user.userType,
      deletedAt: user.deletedAt?.toISOString(),
      created: user.createdAt.toISOString(),
      modified: user.modifiedAt.toISOString(),
    })
  }
}
