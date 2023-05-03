<?php

namespace App\Models;

use App\Traits\WithDiffForHumanTimes;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Database\Eloquent\Model as BaseModel;
use Illuminate\Database\Eloquent\Factories\HasFactory;

class Model extends BaseModel
{
    use HasFactory, SoftDeletes, WithDiffForHumanTimes;
}
