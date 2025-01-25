import { Static, Type as T } from '@sinclair/typebox'
import shortUUID from 'short-uuid'
import { Email } from '../../models/email'
import { Nullable } from '../../../lib/json-schema'

export const UserStatuses = {
  Active: 'ACTIVE',
  Deleted: 'DELETED',
} as const
export type UserStatus = (typeof UserStatuses)[keyof typeof UserStatuses]

export const UserTypes = {
  Customer: 'CUSTOMER',
  Admin: 'ADMIN',
  Consultant: 'CONSULTANT',
} as const
export type UserType = (typeof UserTypes)[keyof typeof UserTypes]

export type Password = {
  hash: string
  salt: string
}

export const UserTypeSchema = T.Object({
  id: T.String(),
  name: T.String(),
  email: T.String({ format: 'email' }),
  status: T.Union([
    T.Literal(UserStatuses.Active),
    T.Literal(UserStatuses.Deleted),
  ]),
  createdAt: T.String(),
  modifiedAt: T.String(),
  deletedAt: Nullable(T.String()),
})

export type UserDTO = Static<typeof UserTypeSchema>

export type UserProps = {
  id?: string
  name: string
  email: Email
  password: Password
  userType: UserType
  status?: UserStatus
  createdAt?: Date
  modifiedAt?: Date
  deletedAt?: Date
}

export class User {
  readonly id: string
  readonly createdAt: Date
  readonly password: Password
  name: string
  email: Email
  userType: UserType
  status: UserStatus
  modifiedAt: Date
  deletedAt?: Date

  constructor(props: UserProps) {
    this.id = props.id ?? `user_${shortUUID.generate()}`
    this.name = props.name
    this.email = props.email
    this.userType = props.userType
    this.password = {
      hash: props.password.hash,
      salt: props.password.salt
    }
    this.status = props.status ?? UserStatuses.Active
    this.createdAt = props.createdAt ?? new Date()
    this.modifiedAt = props.modifiedAt ?? new Date()
    this.deletedAt = props.deletedAt
  }

  delete() {
    this.name = '**REDACTED**'
    this.email = new Email(`${this.id}@deleted.com`)
    this.status = UserStatuses.Deleted
    this.modifiedAt = new Date()
    this.deletedAt = new Date()
  }

  update(input: Partial<{ name: string }>) {
    this.name = input.name ?? this.name
    this.modifiedAt = new Date()
  }

  serialize(): UserDTO {
    return {
      id: this.id,
      name: this.name,
      email: this.email.value,
      status: this.status,
      createdAt: this.createdAt.toISOString(),
      modifiedAt: this.modifiedAt.toISOString(),
      deletedAt: this.deletedAt ? this.deletedAt.toISOString() : null,
    }
  }
}
